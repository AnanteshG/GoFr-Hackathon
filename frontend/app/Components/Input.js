// components/InputField.js
'use client'
import React, { useState } from 'react';

const InputField = ({ onSubmit }) => {
  const [input, setInput] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    if (input.trim() !== '') {
      onSubmit(input);
      setInput('');
    }
  };

  return (
    <form onSubmit={handleSubmit} className="mt-4">
      <input
        type="text"
        value={input}
        onChange={(e) => setInput(e.target.value)}
        className="w-full p-3 border border-gray-300 rounded-md"
        placeholder="Type your message here..."
      />
      <button type="submit" className="w-full mt-2 bg-green-600 text-white p-3 rounded-md">
        Send
      </button>
    </form>
  );
};

export default InputField;
