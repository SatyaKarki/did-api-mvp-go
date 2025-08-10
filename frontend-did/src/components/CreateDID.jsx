import React, { useState } from 'react';
import axios from 'axios';

export default function CreateDID() {
  const [result, setResult] = useState(null);
  const [loading, setLoading] = useState(false);

  const handleCreate = async () => {
    setLoading(true);
    setResult(null);
    try {
      const res = await axios.post('/v1/api/did/create');
      setResult(res.data);
    } catch (err) {
      setResult({ error: err.response?.data || err.message });
    }
    setLoading(false);
  };

  return (
    <section>
      <h2>Create DID</h2>
      <button onClick={handleCreate} disabled={loading} style={{ minWidth: 140 }}>
        {loading ? 'Creating...' : 'Create DID'}
      </button>
      {result && (
        <pre>
          {JSON.stringify(result, null, 2)}
        </pre>
      )}
    </section>
  );
}
