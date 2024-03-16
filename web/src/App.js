import React, { useEffect, useState } from 'react';
import './App.css';

function App() {
  const [users, setUsers] = useState([]);
  const userUrl = process.env.REACT_APP_USER_SERVICE_URL;
  const musicUrl = process.env.REACT_APP_MUSIC_SERVICE_URL;
  const playlistUrl = process.env.REACT_APP_PLAYLIST_SERVICE_URL;
  useEffect(() => {
    fetch(userUrl+'/users')
      .then(response => response.json())
      .then(data => setUsers(data));
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <p>Users</p>
        <ul>
          {users.map(user => (
            <li key={user.id}>{user.name}</li>
          ))}
        </ul>
      </header>
    </div>
  );
}

export default App;
