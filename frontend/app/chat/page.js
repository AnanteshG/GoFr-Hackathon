'use client'

import React, { useState } from 'react';
import InputField from '../Components/Input';
import ServerPrompt from '../Components/Server';
import UserPrompt from './../Components/user-box';
import Title from '../Components/Title';

const ChatPage = () => {
  const [messages, setMessages] = useState([]);

  // Function to handle user input
  const handleUserInput = (userMessage) => {
    if (!userMessage.trim()) return; // Ignore empty messages
    setMessages((prevMessages) => [
      ...prevMessages,
      { type: 'user', text: userMessage }
    ]);

    // Simulating server response
    setTimeout(() => {
      const botResponse = `Bot: I received your message: "${userMessage}"`;
      setMessages((prevMessages) => [
        ...prevMessages,
        { type: 'server', text: botResponse }
      ]);
    }, 1000);
  };

  return (
    <div className="min-h-screen bg-gradient-to-b from-green-500 to-green-600 flex flex-col items-center py-10">
      {/* Page Title */}
      <Title />

      {/* Chat Messages Section */}
      <div className="w-full max-w-lg bg-white p-6 rounded-lg shadow-lg flex flex-col space-y-4 overflow-y-auto max-h-[65vh]">
        {messages.map((message, index) =>
          message.type === 'user' ? (
            <UserPrompt key={index} message={message.text} />
          ) : (
            <ServerPrompt key={index} message={message.text} />
          )
        )}
      </div>

      {/* Input Field for User Message */}
      <div className="w-full max-w-lg mt-4">
        <InputField onSubmit={handleUserInput} />
      </div>
    </div>
  );
};

export default ChatPage;



