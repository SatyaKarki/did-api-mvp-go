
import React from 'react';
import './App.css';
import HealthCheck from './components/HealthCheck';
import CreateDID from './components/CreateDID';
import ResolveDID from './components/ResolveDID';
import AddKeyToDID from './components/AddKeyToDID';
import SignMessage from './components/SignMessage';
import VerifyMessage from './components/VerifyMessage';

export default function App() {
  return (
    <div className="app-container">
      <h1>DID API Frontend</h1>
      <HealthCheck />
      <hr />
      <CreateDID />
      <hr />
      <ResolveDID />
      <hr />
      <AddKeyToDID />
      <hr />
      <SignMessage />
      <hr />
      <VerifyMessage />
    </div>
  );
}
