// Register.js

import React, { useState } from 'react';
import { TextField, Button } from '@mui/material';

const Register = () => {
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleFirstNameChange = (e) => {
    setFirstName(e.target.value);
  };

  const handleLastNameChange = (e) => {
    setLastName(e.target.value);
  };

  const handleEmailChange = (e) => {
    setEmail(e.target.value);
  };

  const handlePasswordChange = (e) => {
    setPassword(e.target.value);
  };

  const handleRegister = () => {
    const user = {
      firstName,
      lastName,
      email,
      password
    };

    fetch('http://localhost:8000/api/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(user)
    })
      .then((response) => response.json())
      .then((data) => {
        // Gérer la réponse de l'API ici
        console.log(data);
      })
      .catch((error) => {
        // Gérer les erreurs de la requête ici
        console.error(error);
      });
  };

  return (
    <div>
      <TextField label="Prénom" value={firstName} onChange={handleFirstNameChange} />
      <TextField label="Nom" value={lastName} onChange={handleLastNameChange} />
      <TextField label="Email" value={email} onChange={handleEmailChange} />
      <TextField label="Mot de passe" value={password} onChange={handlePasswordChange} type="password" />
      <Button variant="contained" color="primary" onClick={handleRegister}>
        S'inscrire
      </Button>
    </div>
  );
};

export default Register;
