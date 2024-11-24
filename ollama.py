import os
from flask import Flask, request, jsonify
from langchain.document_loaders import TextLoader
from langchain_community.embeddings import OllamaEmbeddings
from langchain_text_splitters import RecursiveCharacterTextSplitter
from langchain_community.vectorstores import Chroma
from langchain.prompts import PromptTemplate, ChatPromptTemplate
from langchain_core.output_parsers import StrOutputParser
from langchain_community.chat_models import ChatOllama
from langchain_core.runnables import RunnablePassthrough
from langchain.retrievers.multi_query import MultiQueryRetriever

# Initialize Flask app
app = Flask(__name__)

# Local PDF or text file setup
local_path = "data.txt"  # You can change this path to your file

# Load documents if local path is provided
if local_path:
    loader = TextLoader(local_path, encoding="latin1")
    data = loader.load()

# Split and chunk the document
text_splitter = RecursiveCharacterTextSplitter(chunk_size=1000, chunk_overlap=100)
chunks = text_splitter.split_documents(data)

# Setup for Chroma vector store
current_dir = os.getcwd()
persistent_directory = os.path.join(current_dir, "db", "chroma_db_for_MotorAct")
embedding_function = OllamaEmbeddings(model="nomic-embed-text", show_progress=True)

# Check if vector store already exists
if os.path.exists(persistent_directory):
    vector_db = Chroma(
        persist_directory=persistent_directory,
        embedding_function=embedding_function,
        collection_name="local-rag"
    )
    print("Loaded existing Chroma vector store.")
else:
    vector_db = Chroma.from_documents(
        documents=chunks,
        embedding=embedding_function,
        collection_name="local-rag",
        persist_directory=persistent_directory
    )
    vector_db.persist()

# LLM from Ollama setup
local_model = "llama3.2"  # Changed from llama3.2 to llama2 as it's more commonly available
llm = ChatOllama(model=local_model)

# Query prompt for multi-query retrieval
QUERY_PROMPT = PromptTemplate(
    input_variables=["question"],
    template="""You are a GitHub Repository who answers questions taking relevance from the data present in the vector database. By generating multiple perspectives on the user question, your
    goal is to help the user overcome some of the limitations of the distance-based
    similarity search. Provide these alternative questions separated by newlines.
    Original question: {question}""",
)

# Set up retriever using MultiQueryRetriever
retriever = MultiQueryRetriever.from_llm(
    vector_db.as_retriever(),
    llm,
    prompt=QUERY_PROMPT
)

# Set up the RAG chain
def format_docs(docs):
    return "\n\n".join(doc.page_content for doc in docs)

# Chat prompt for answering based on legal context
template = """You are a GitHub Repository. Answer the question based ONLY on the following context:
{context}
Question: {question}
"""
prompt = ChatPromptTemplate.from_template(template)

chain = (
    {
        "context": retriever | format_docs,
        "question": RunnablePassthrough()
    }
    | prompt
    | llm
    | StrOutputParser()
)

# Flask route to handle POST request for a question
@app.route('/api/chat', methods=['POST'])
def ask():
    try:
        # Get the JSON data from the request
        data = request.get_json(force=True)  # Added force=True to handle content-type issues
        if not data:
            return jsonify({"error": "No JSON data provided"}), 400

        # Get the question - try both "Question" and "question" keys
        question = data.get('Question') or data.get('question')
        print(f"Received question: {question}")
        if question == "":
            return jsonify({"error": "No question found in request. Please provide either 'question' or 'Question' in JSON"}), 400

    

        # Use the chain to get the response
        response = chain.invoke(question)
        
        return jsonify({"response": response})
    
    except Exception as e:
        print(f"Error occurred: {str(e)}")
        return jsonify({"error": f"An error occurred: {str(e)}"}), 500

# Start the Flask server
if __name__ == '__main__':
    app.run(host='localhost', port=8000)  # Changed port to 5000 and host to 0.0.0.0 for better accessibility