import React, { useState } from 'react';
import axios from 'axios';

export default function ResolveDID() {
  const [did, setDid] = useState('');
  const [result, setResult] = useState(null);
  const [loading, setLoading] = useState(false);

  const handleResolve = async () => {
    setLoading(true);
    setResult(null);
    try {
      const res = await axios.get(`/v1/api/did/resolve/${did}`);
      setResult(res.data);
    } catch (err) {
      setResult({ error: err.response?.data || err.message });
    }
    setLoading(false);
  };

  return (
    <section>
      <h2>Resolve DID</h2>
      <input
        type="text"
        value={did}
        onChange={e => setDid(e.target.value)}
        placeholder="Enter DID (e.g. did:key:...)"
      />
      <button onClick={handleResolve} disabled={loading || !did} style={{ minWidth: 120 }}>
        {loading ? 'Resolving...' : 'Resolve'}
      </button>
      {result && (
        <pre>
          {JSON.stringify(result, null, 2)}
        </pre>
      )}
    </section>
  );
}
