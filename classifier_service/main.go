package main

import (
	"context"
	"log"
	"net/http"

	pb "classification_project/config/stub" // замените на путь к вашему сгенерированному пакету protobuf
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Request struct {
	ID          string `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type Response struct {
	ID  string `json:"id"`
	Tag string `json:"tag"`
}

func main() {
	router := gin.Default()
	router.POST("/get/tag", handleRequest)
	log.Println("Starting server on :8080")
	log.Fatal(router.Run(":8080"))
}

func handleRequest(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag, err := classifyDescription(req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to classify description"})
		return
	}

	resp := Response{
		ID:  req.ID,
		Tag: tag,
	}

	c.JSON(http.StatusOK, resp)
}

func classifyDescription(description string) (string, error) {
	// Устанавливаем соединение с gRPC сервером Python
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := pb.NewClassifierClient(conn)
	resp, err := client.Classify(context.Background(), &pb.ClassifyRequest{Description: description})
	if err != nil {
		return "", err
	}

	return resp.Tag, nil
}
