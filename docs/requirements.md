# ForgeStore - System Requirements

## Functional Requirements

1. Upload files through HTTP.
2. Split large files into chunks.
3. Calculate SHA-256 checksum for each chunk.
4. Store chunks on multiple storage nodes.
5. Replicate chunks across different nodes.
6. Store file metadata in PostgreSQL.
7. Download and reconstruct original files.
8. Verify chunk integrity during retrieval.
9. Detect unavailable storage nodes.
10. Delete files and their associated chunks.

## Non-Functional Requirements

1. Memory-efficient streaming.
2. Support concurrent uploads.
3. Data durability through replication.
4. Fault tolerance against individual node failures.
5. Data integrity verification.
6. Clear service boundaries.
7. Observable system behavior.

## Initial Configuration

- Language: Go
- Database: PostgreSQL
- Chunk Size: 10 MB
- Storage Nodes: 4
- Replication Factor: 2
- Communication: HTTP
- Checksum: SHA-256

## Initial Limitations

- No automatic leader election.
- No consensus protocol.
- No distributed transactions.
- No multi-region deployment.
- No automatic rebalancing in the first MVP.