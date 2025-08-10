import React, { useState } from 'react';
import axios from 'axios';

export default function AddKeyToDID() {
  const [did, setDid] = useState('');
  const [publicKey, setPublicKey] = useState('');
  const [type, setType] = useState('Ed25519VerificationKey2020');
  const [keyId, setKeyId] = useState('keys-2');
  const [result, setResult] = useState(null);
  const [loading, setLoading] = useState(false);

  const handleAddKey = async () => {
    setLoading(true);
    setResult(null);
    try {
      const res = await axios.post(`/v1/api/did/add-key/${did}`, {
        publicKey,
        type,
        keyId
      });
      setResult(res.data);
    } catch (err) {
      setResult({ error: err.response?.data || err.message });
    }
    setLoading(false);
  };

  return (
    <section>
      <h2>Add Key to DID</h2>
      <input
        type="text"
        value={did}
        onChange={e => setDid(e.target.value)}
        placeholder="DID (e.g. did:key:...)"
      />
      <input
        type="text"
        value={publicKey}
        onChange={e => setPublicKey(e.target.value)}
        placeholder="Public Key (Base58)"
      />
      <input
        type="text"
        value={type}
        onChange={e => setType(e.target.value)}
        placeholder="Type"
      />
      <input
        type="text"
        value={keyId}
        onChange={e => setKeyId(e.target.value)}
        placeholder="Key ID"
      />
      <button onClick={handleAddKey} disabled={loading || !did || !publicKey || !type || !keyId} style={{ minWidth: 120 }}>
        {loading ? 'Adding...' : 'Add Key'}
      </button>
      {result && (
        <pre>
          {JSON.stringify(result, null, 2)}
        </pre>
      )}
    </section>
  );
}
