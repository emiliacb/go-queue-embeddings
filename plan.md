## Outline

- Receives the PDF and returns a UUID process id
- The request is a POST request with the following fields:
  - pdf: the PDF file (multipart/form-data)
  - chunk_strategy: the strategy to divide the PDF into chunks
- The response is a JSON with the following fields:
  - id: UUID process id
  - status: the status of the process
  - progress: the progress of the process
- Divides the PDF into chunks
- Sends the chunk to the embedding service
- Saves the result as a JSON file with the name <process_id>.json
- The result
- The JSON file has the following fields:
  - id: UUID process id
  - status: the status of the process
  - progress: the progress of the process
  - data: the chunks of the result
  - metadata: {
    - chunk_size: the size of the chunks in tokens
    - embedding_model: the embedding model that was used
  }
- Each chunk has the following fields:
  - id: UUID chunk id
  - text: the text of the chunk
  - embedding: the embedding of the chunk
  - metadata: {
    - chunk_size: the size of the chunks in tokens
    - embedding_model: the embedding model that was used
  }

## POC VERSION

TODO:
- [X] setup repo and write the README.md
- [X] set up basic server with GIN
- [X] simple /embed endpoint with Ollama
- [X] interfaces: 
    - [ ] queue: 
        - send(job: Job)
        - receive(id: string) -> Job
    - [X] embeddingAdapter:
        - embed(text: string, config: EmbedConfig) -> []float32
        - config:
            - model: string
            - api_key: string
    - [X] chunkStrategy:
        - chunk(text: string, config: ChunkConfig) -> []string
        - config:
            - chunk_size: int
            - chunk_overlap: int
    - [X] documentStorageAdapter:
        - upsert(id: string, status: string, progress: int, data, metadata)
        - get(id: string) -> ([]byte, error)
- [ ] docker config and deploy to Hugging Face ASAP to avoid having unknown issues with the server later.
- [ ] basic queue and return the process id
- [ ] implement /process/<process_id> endpoint to return the JSON file
- [ ] create a simple frontend to upload the PDF and see the process with short polling (HTMX now, later we can use Next.js)
- [ ] save the json file
- [ ] implement the file upload and chunking
- [ ] Add structured logs
- [ ] Add tests