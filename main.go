package main

import (
	"github.com/hantmac/langchaingo-ollama-rag/rag"
	"github.com/hantmac/langchaingo-ollama-rag/rag/logger"
)

func init() {
	logger.InitLogger("debug")
}

func main() {
	rag.Execute()
}
