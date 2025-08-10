import React, { useState } from 'react';
import axios from 'axios';

export default function HealthCheck() {
  const [status, setStatus] = useState('');
  const checkHealth = async () => {
    try {
      const res = await axios.get('/ping');
      setStatus(res.data);
    } catch (err) {
      setStatus('Error: ' + (err.response?.data || err.message));
    }
  };
  return (
    <section>
      <h2>Health Check</h2>
      <button onClick={checkHealth} style={{ minWidth: 120 }}>Check API</button>
      {status && (
        <div style={{ marginTop: 12, fontWeight: 500, color: status === 'pong' ? '#22c55e' : '#ef4444' }}>
          Status: <b>{status}</b>
        </div>
      )}
    </section>
  );
}
