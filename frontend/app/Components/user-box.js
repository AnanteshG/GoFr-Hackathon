// components/UserPrompt.js
import React from 'react';

const UserPrompt = ({ message }) => (
  <div className="bg-blue-100 p-4 rounded-lg mb-2">
    <strong>User:</strong> {message}
  </div>
);

export default UserPrompt;
