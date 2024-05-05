package controller

import (
	"app/utils"
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	openai "github.com/sashabaranov/go-openai"
)

func RecommendationAI(c echo.Context) error {
	userInput := c.FormValue("Gol_Darah")

	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	// Meminta rekomendasi golongan darah yang cocok kepada AI
	client := openai.NewClient(os.Getenv("AI_TOKEN"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Anda cukup memberitahu golongan darah mana yang dapat diterima oleh golongan darah: " + userInput,
				},
			},
		},
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Gagal menghubungi AI"))
	}

	// Mengambil jawaban dari AI
	recommendedBloodType := resp.Choices[0].Message.Content

	return c.JSON(http.StatusOK, utils.SuccessResponse("Rekomendasi: ", recommendedBloodType))
}
