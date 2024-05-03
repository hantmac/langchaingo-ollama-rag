# Ollama RAG
Project for adding RAG to ollama models using :
- [langchaingo](https://github.com/tmc/langchaingo)
- [qdrant vector database](https://github.com/qdrant/qdrant)
- [nomic-embed-text](https://ollama.com/library/nomic-embed-text)
- [qwen1.8B](https://github.com/QwenLM/Qwen1.5)

## Installation :
Download packages :
```bash
go mod download
```
Pull Ollama models :
```
ollama pull nomic-embed-text
ollama pull qwen1.8B
```

Install qdrant :
```bash
docker pull qdrant/qdrant
docker run -p 6333:6333 qdrant/qdrant
```
Create Collection for qdrant
use any http client for make a PUT request like example blow for creating a Collection

```bash
curl -X PUT http://localhost:6333/collections/romeo \
  -H 'Content-Type: application/json' \
  --data-raw '{
    "vectors": {
      "size": 768,
      "distance": "Dot"
    }
  }'
```

Delete a collection
```bash
curl --location --request DELETE 'http://localhost:6333/collections/romeo'
```

## Using

put you texts in text.txt and run :
``` go run main.go getanswer
```
then ask anything from the text you provided


## Reference
Thanks to https://github.com/shayanfarzi/ollama-rag
