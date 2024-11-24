// 'use client';

// import React, { useState } from 'react';
// import { useRouter } from 'next/navigation';

// const Page = () => {
//   const [url, setUrl] = useState('');
//   const [username, setUsername] = useState('');
//   const [repoName, setRepoName] = useState('');
//   const [githubAccessToken, setGithubAccessToken] = useState('');
//   const [error, setError] = useState('');
//   const router = useRouter();

//   const handleSubmit = async (e) => {
//     e.preventDefault();
//     setError('');

//     const trimmedToken = githubAccessToken.trim();

//     // Validate GitHub access token format
//     if (!trimmedToken.startsWith('ghp_') && !trimmedToken.startsWith('github_pat_')) {
//       setError('Invalid GitHub Access Token. Token should start with ghp_ or github_pat_');
//       return;
//     }

//     if (trimmedToken.length < 40) {
//       setError('Invalid GitHub Access Token length. Token should be at least 40 characters long');
//       return;
//     }

//     try {
//       // Check if the repository exists
//       const response = await fetch(`https://api.github.com/repos/${username}/${repoName}`, {
//         headers: {
//           Authorization: `Bearer ${trimmedToken}`,
//         },
//       });

//       if (response.status === 404) {
//         setError('Repository not found. Please check the username and repository name.');
//         return;
//       } else if (!response.ok) {
//         setError('An error occurred while validating the repository. Please try again.');
//         return;
//       }

//       // If successful, redirect to the chat page
//       // console.log({ url, username, repoName });
//       else{
//       console.log({ url, username, repoName });
//       router.push('/Chat');
//       }
//     } catch (err) {
//       setError('Failed to validate the repository. Please check your internet connection.');
//     }
//   };

//   return (
//     <div className="w-full h-full bg-green-300 flex justify-center items-center p-6">
//       <div className="bg-white p-8 rounded-lg shadow-lg w-full max-w-lg">
//         <h1 className="text-2xl font-bold mb-6 text-center">Go Get Git</h1>
//         <form onSubmit={handleSubmit}>
//           <div className="mb-4">
//             <label htmlFor="url" className="block text-sm font-semibold text-gray-700">
//               Github URL <span className="text-red-500">*</span>
//             </label>
//             <input
//               type="url"
//               placeholder="Enter Your Github URL"
//               id="url"
//               value={url}
//               onChange={(e) => setUrl(e.target.value)}
//               required
//               className="mt-2 w-full p-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
//             />
//           </div>

//           <div className="mb-4">
//             <label htmlFor="username" className="block text-sm font-semibold text-gray-700">
//               Github Username <span className="text-red-500">*</span>
//             </label>
//             <input
//               type="text"
//               placeholder="Enter Your Github Username"
//               id="username"
//               value={username}
//               onChange={(e) => setUsername(e.target.value)}
//               required
//               className="mt-2 w-full p-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
//             />
//           </div>

//           <div className="mb-4">
//             <label htmlFor="repoName" className="block text-sm font-semibold text-gray-700">
//               Repo Name <span className="text-red-500">*</span>
//             </label>
//             <input
//               type="text"
//               placeholder="Enter Repo Name"
//               id="repoName"
//               value={repoName}
//               onChange={(e) => setRepoName(e.target.value)}
//               required
//               className="mt-2 w-full p-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500"
//             />
//           </div>

//           <div className="mb-4">
//             <label htmlFor="githubAccessToken" className="block text-sm font-semibold text-gray-700">
//               GitHub Access Token <span className="text-red-500">*</span>
//             </label>
//             <input
//               type="password"
//               id="githubAccessToken"
//               value={githubAccessToken}
//               onChange={(e) => setGithubAccessToken(e.target.value)}
//               required
//               className="mt-2 w-full p-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500"
//             />
//           </div>

//           {error && (
//             <div className="text-red-500 text-center mb-4">
//               {error}
//             </div>
//           )}

//           <div className="text-center">
//             <button
//               type="submit" 
//               className="w-full bg-green-500 text-white py-3 rounded-md font-semibold hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-green-500"
//             >
//               Submit
//             </button>
//           </div>
//         </form>
//       </div>
//     </div>
//   );
// };

// export default Page;


'use client';

import React, { useState } from 'react';
import { useRouter } from 'next/navigation';

const Page = () => {
  const [url, setUrl] = useState('');
  const [username, setUsername] = useState('');
  const [repoName, setRepoName] = useState('');
  const [githubAccessToken, setGithubAccessToken] = useState('');
  const [error, setError] = useState('');
  const router = useRouter();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');

    const trimmedToken = githubAccessToken.trim();

    // Validate GitHub access token format
    if (!trimmedToken.startsWith('ghp_') && !trimmedToken.startsWith('github_pat_')) {
      setError('Invalid GitHub Access Token. Token should start with ghp_ or github_pat_');
      return;
    }

    if (trimmedToken.length < 40) {
      setError('Invalid GitHub Access Token length. Token should be at least 40 characters long');
      return;
    }

    try {
      // Check if the repository exists
      const response = await fetch(`https://api.github.com/repos/${username}/${repoName}`, {
        headers: {
          Authorization: `Bearer ${trimmedToken}`,
        },
      });

      if (response.status === 404) {
        setError('Repository not found. Please check the username and repository name.');
        return;
      } else if (!response.ok) {
        setError('An error occurred while validating the repository. Please try again.');
        return;
      }

      // If successful, redirect to the chat page
      console.log({ url, username, repoName });
      router.push('/chat');
    } catch (err) {
      setError('Failed to validate the repository. Please check your internet connection.');
    }
  };

  // Extract username and repo from URL
  const handleUrlChange = (e) => {
    const newUrl = e.target.value;
    setUrl(newUrl);
    const regex = /github\.com\/([^/]+)\/([^/]+)/;
    const match = newUrl.match(regex);
    if (match) {
      setUsername(match[1]);
      setRepoName(match[2]);
    }
  };

  return (
    <div className="w-full h-full bg-green-300 flex justify-center items-center p-6">
      <div className="bg-white p-8 rounded-lg shadow-lg w-full max-w-lg">
        <h1 className="text-2xl font-bold mb-6 text-center">Go Get Git</h1>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label htmlFor="url" className="block text-sm font-semibold text-gray-700">
              Github URL <span className="text-red-500">*</span>
            </label>
            <input
              type="url"
              placeholder="Enter Your Github URL"
              id="url"
              value={url}
              onChange={handleUrlChange}
              required
              className="mt-2 w-full p-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500"
            />
          </div>

          <div className="mb-4">
            <label htmlFor="username" className="block text-sm font-semibold text-gray-700">
              Github Username <span className="text-red-500">*</span>
            </label>
            <input
              type="text"
              placeholder="Enter Your Github Username"
              id="username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
              className="mt-2 w-full p-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500"
            />
          </div>

          <div className="mb-4">
            <label htmlFor="repoName" className="block text-sm font-semibold text-gray-700">
              Repo Name <span className="text-red-500">*</span>
            </label>
            <input
              type="text"
              placeholder="Enter Repo Name"
              id="repoName"
              value={repoName}
              onChange={(e) => setRepoName(e.target.value)}
              required
              className="mt-2 w-full p-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500"
            />
          </div>

          <div className="mb-4">
            <label htmlFor="githubAccessToken" className="block text-sm font-semibold text-gray-700">
              GitHub Access Token <span className="text-red-500">*</span>
            </label>
            <input
              type="password"
              id="githubAccessToken"
              value={githubAccessToken}
              onChange={(e) => setGithubAccessToken(e.target.value)}
              required
              className="mt-2 w-full p-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-green-500"
            />
          </div>

          {error && (
            <div className="text-red-500 text-center mb-4">
              {error}
            </div>
          )}

          <div className="text-center">
            <button
              type="submit" 
              className="w-full bg-green-500 text-white py-3 rounded-md font-semibold hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-green-500"
            >
              Submit
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Page;



