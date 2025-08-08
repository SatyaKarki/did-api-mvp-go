// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract DIDRegistry {
    mapping(string => string) private didDocuments;

    event DIDDocumentSet(string did, string document);

    function setDIDDocument(string memory did, string memory document) public {
        didDocuments[did] = document;
        emit DIDDocumentSet(did, document);
    }

    function getDIDDocument(string memory did) public view returns (string memory) {
        return didDocuments[did];
    }
}
