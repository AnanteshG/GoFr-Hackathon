// components/ServerPrompt.js
import React from 'react';

const ServerPrompt = ({ message }) => (
  <div className="bg-gray-100 p-4 rounded-lg mb-2">
    <strong>Bot:</strong> {message}
  </div>
);

export default ServerPrompt;
