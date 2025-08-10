import React, { useState } from 'react';
import axios from 'axios';

export default function VerifyMessage() {
  const [message, setMessage] = useState('');
  const [signature, setSignature] = useState('');
  const [publicKey, setPublicKey] = useState('');
  const [result, setResult] = useState(null);
  const [loading, setLoading] = useState(false);

  const handleVerify = async () => {
    setLoading(true);
    setResult(null);
    try {
      const params = new URLSearchParams();
      params.append('message', message);
      params.append('signature', signature);
      params.append('publicKey', publicKey);
      const res = await axios.post('/v1/api/did/verify', params);
      setResult(res.data);
    } catch (err) {
      setResult({ error: err.response?.data || err.message });
    }
    setLoading(false);
  };

  return (
    <section>
      <h2>Verify Message</h2>
      <input
        type="text"
        value={message}
        onChange={e => setMessage(e.target.value)}
        placeholder="Message"
      />
      <input
        type="text"
        value={signature}
        onChange={e => setSignature(e.target.value)}
        placeholder="Signature (Base58)"
      />
      <input
        type="text"
        value={publicKey}
        onChange={e => setPublicKey(e.target.value)}
        placeholder="Public Key (Base58)"
      />
      <button onClick={handleVerify} disabled={loading || !message || !signature || !publicKey} style={{ minWidth: 120 }}>
        {loading ? 'Verifying...' : 'Verify'}
      </button>
      {result && (
        <pre>
          {JSON.stringify(result, null, 2)}
        </pre>
      )}
    </section>
  );
}
