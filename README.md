# go-queue-embeddings
A minimal Go service that queues long-running embedding tasks. 

**Goals:**
- Showcase concurrency patterns in Go using worker queues
- Provide a working pipeline for document processing and embedding
- Focus on modular interfaces for testing, mocking, and future extensibility

### Key Decisions
- Gin because is popular and easy to use
- Using freely an hexagonal architecture approach to ensure extensibility, specially decoupling the logic from the embedding provider and the storage
- Saving as the process JSON in a temp folder for this POC but the code is expandable to save in a database or other storage in the future

### API Flow
1. **PDF Upload**
   - Route: `/upload`
   - Receives PDF via POST request
   - Returns UUID process ID for tracking
   - Request fields:
     - `pdf`: PDF file (multipart/form-data)
     - `chunk_strategy`: Strategy for PDF text chunking

2. **Response Format**
   ```json
   {
     "id": "uuid-process-id",
     "status": "processing",
     "progress": 25
   }
   ```

3. **Processing Pipeline**
   - PDF is divided into chunks based on strategy
   - Each chunk is sent to embedding service
   - Results saved as JSON file (`<process_id>.json`)


4. **Output route**
   - Route: `/process/<process_id>`
   - Returns the JSON file with the status and, if completed, the results

### Data Model
 
1. **Process JSON**
   ```json
   {
     "id": "uuid-process-id",
     "status": "processing|completed|failed",
     "progress": 75,
     "data": [
       {
         "id": "uuid-chunk-id",
         "text": "chunk text content",
         "embedding": [0.1, 0.2, ...],
         "metadata": {
           "chunk_size": 512,
           "embedding_model": "model-name"
         }
       }
     ],
     "metadata": {
       "chunk_size": 512,
       "embedding_model": "model-name"
     }
   }
   ```
