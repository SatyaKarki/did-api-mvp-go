import React, { useState } from 'react';
import axios from 'axios';

export default function SignMessage() {
  const [message, setMessage] = useState('');
  const [privateKey, setPrivateKey] = useState('');
  const [result, setResult] = useState(null);
  const [loading, setLoading] = useState(false);

  const handleSign = async () => {
    setLoading(true);
    setResult(null);
    try {
      const params = new URLSearchParams();
      params.append('message', message);
      params.append('privateKey', privateKey);
      const res = await axios.post('/v1/api/did/sign', params);
      setResult(res.data);
    } catch (err) {
      setResult({ error: err.response?.data || err.message });
    }
    setLoading(false);
  };

  return (
    <section>
      <h2>Sign Message</h2>
      <input
        type="text"
        value={message}
        onChange={e => setMessage(e.target.value)}
        placeholder="Message to sign"
      />
      <input
        type="text"
        value={privateKey}
        onChange={e => setPrivateKey(e.target.value)}
        placeholder="Private Key (Base58)"
      />
      <button onClick={handleSign} disabled={loading || !message || !privateKey} style={{ minWidth: 120 }}>
        {loading ? 'Signing...' : 'Sign'}
      </button>
      {result && (
        <pre>
          {JSON.stringify(result, null, 2)}
        </pre>
      )}
    </section>
  );
}
