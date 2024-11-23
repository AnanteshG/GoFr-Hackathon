from flask import Flask, request, jsonify
from langchain_ollama import ChatOllama
from langchain.prompts import PromptTemplate
from langchain_core.output_parsers import StrOutputParser
from langchain_community.embeddings import OllamaEmbeddings
from langchain_text_splitters import RecursiveCharacterTextSplitter
from langchain_community.vectorstores import Chroma
import os

# Initialize Flask app
app = Flask(__name__)

# Initialize the Ollama model
local_model = "llama3.2"
llm = ChatOllama(model=local_model)

# Vector store and embeddings setup
current_dir = os.getcwd()
persistent_directory = os.path.join(current_dir, "db", "chroma_db_for_GitHub")
embedding_function = OllamaEmbeddings(model="nomic-embed-text", show_progress=True)

# Route to handle prompts
@app.route('/ask', methods=['POST'])
def ask():
    # Get the user input from the POST request
    data = request.get_json()
    question = data.get("question")

    if not question:
        return jsonify({"error": "No question provided"}), 400

    # Assuming vector DB exists, otherwise, you would create it (similar to your notebook)
    if os.path.exists(persistent_directory):
        vector_db = Chroma(
            persist_directory=persistent_directory,
            embedding_function=embedding_function,
            collection_name="local-rag"
        )
    else:
        return jsonify({"error": "Vector store not found. Please load documents first."}), 404

    # Create a retrieval prompt
    query_prompt = PromptTemplate(
        input_variables=["question"],
        template="""You are a GitHub Repository who answers questions taking
                    relevance from the data present in the vector database. By generating
                    multiple perspectives on the user question, your goal is to help the user 
                    overcome some of the limitations of the distance-based similarity search. 
                    Provide these alternative questions separated by newlines.
                    Original question: {question}""",
    )

    # Perform retrieval and process with LLM (Ollama)
    retriever = vector_db.as_retriever()
    response = retriever.query(question)
    return jsonify({"response": response})

# Start the Flask server
if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
