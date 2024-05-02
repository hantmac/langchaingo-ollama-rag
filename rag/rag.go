package rag

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/hantmac/langchaingo-ollama-rag/rag/logger"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "langchaingo-ollama",
	Short: "学习基于langchaingo构建的问答系统",
	Long:  `学习基于langchaingo构建的问答系统`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// ========
	rootCmd.AddCommand(GetAnwserCmd)
	GetAnwserCmd.Flags().IntP("topk", "t", 5, "召回数据的数量，默认为5")
}

func FileToChunksCmd() {
	filepath := "test.txt"
	chunkSize := 5
	chunkOverlap := 2

	docs, err := TextToChunks(filepath, chunkSize, chunkOverlap)
	if err != nil {
		logger.Error("转换文件为块儿失败，错误信息: %v", err)
	}
	logger.Info("转换文件为块儿成功，块儿数量: ", len(docs))
	for _, v := range docs {
		fmt.Printf("🗂 块儿内容==> %v\n", v.PageContent)
	}
}

func EmbeddingCmd() {
	filepath := "test.txt"
	chunkSize := 5
	chunkOverlap := 2
	docs, err := TextToChunks(filepath, chunkSize, chunkOverlap)
	if err != nil {
		logger.Error("转换文件为块儿失败，错误信息: %v", err)
	}
	err = storeDocs(docs, getStore())
	if err != nil {
		logger.Error("转换块儿为向量失败，错误信息: %v", err)
	} else {
		logger.Info("转换块儿为向量成功")
	}
}

var GetAnwserCmd = &cobra.Command{
	Use:   "getanswer",
	Short: "获取回答",
	Run: func(cmd *cobra.Command, args []string) {
		topk, _ := cmd.Flags().GetInt("topk")
		FileToChunksCmd()
		EmbeddingCmd()

		prompt, err := GetUserInput("请输入你的问题")
		if err != nil {
			logger.Error("获取用户输入失败，错误信息: %v", err)
		}
		rst, err := useRetriaver(getStore(), prompt, topk)
		if err != nil {
			logger.Error("检索文档失败，错误信息: %v", err)
		}
		answer, err := GetAnswer(context.Background(), getOllamaQwen(), rst, prompt)
		if err != nil {
			logger.Error("获取回答失败，错误信息: %v", err)
		} else {
			fmt.Printf("🗂 回答==> %s\n\n", answer)
		}
	},
}
