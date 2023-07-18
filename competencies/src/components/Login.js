import React, { useState } from 'react';
import { Link } from "react-router-dom";


const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleEmailChange = (e) => {
    setEmail(e.target.value);
  };

  const handlePasswordChange = (e) => {
    setPassword(e.target.value);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    // Effectuer une requête HTTP pour connecter l'utilisateur avec l'email et le mot de passe
    // Utilisez la bibliothèque de votre choix, telle que 'axios' ou la fonction native 'fetch'
  };

  return (
    <div>
      <h2>Page de connexion</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="email">Email:</label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={handleEmailChange}
          />
        </div>
        <div>
          <label htmlFor="password">Mot de passe:</label>
          <input
            type="password"
            id="password"
            value={password}
            onChange={handlePasswordChange}
          />
        </div>
        <button type="submit">Se connecter</button>
      </form>
      <p>
          Vous n'avez pas de compte ? <Link to="/register">Inscrivez-vous ici</Link>.
        </p>
    </div>
  );
};

export default Login;
