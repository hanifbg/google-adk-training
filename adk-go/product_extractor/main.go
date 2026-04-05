package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/cmd/launcher/full"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/genai"
)

// ProductInfo defines the structured output schema for the model
type ProductInfo struct {
	ProductName string  `json:"product_name" doc:"The full name of the product"`
	Price       float64 `json:"price" doc:"The price in USD"`
	Storage     string  `json:"storage" doc:"Storage capacity (e.g., '256GB')"`
	Color       string  `json:"color" doc:"Product color if mentioned"`
}

func main() {
	ctx := context.Background()

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Note: .env file not found, using system environment variables.")
	}

	// Initialize Model (using Vertex AI via environment variables)
	model, err := gemini.NewModel(ctx, "gemini-2.5-flash", &genai.ClientConfig{
		Backend:  genai.BackendVertexAI,
		Project:  os.Getenv("GOOGLE_CLOUD_PROJECT"),
		Location: os.Getenv("GOOGLE_CLOUD_LOCATION"),
	})
	if err != nil {
		log.Fatalf("Failed to create model: %v", err)
	}

	// Configure the Agent
	llmConfig := llmagent.Config{
		Name:        "product_extractor",
		Model:       model,
		Description: "Extracts product information from user messages and returns structured JSON",
		Instruction: `You are a Product Information Extractor.

Your task:
- Read the user's message about a product.
- Extract: product_name, price, storage, and color (if mentioned).
- Respond ONLY with valid JSON matching the defined schema.

Rules:
- 'price' must be a number (no dollar signs or text).
- 'storage' must include units (GB, TB).
- If color is not mentioned, use "Not specified".
- Output ONLY the JSON, no explanation or markdown text.`,

		// Enforce the structured output
		OutputSchema: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"product_name": {
					Type:        genai.TypeString,
					Description: "The full name of the product",
				},
				"price": {
					Type:        genai.TypeNumber,
					Description: "The price in USD",
				},
				"storage": {
					Type:        genai.TypeString,
					Description: "Storage capacity (e.g., '256GB')",
				},
				"color": {
					Type:        genai.TypeString,
					Description: "Product color if mentioned",
				},
			},
			Required: []string{"product_name", "price", "storage", "color"},
		},
		OutputKey:    "extracted_product",
	}

	agentInstance, err := llmagent.New(llmConfig)
	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	// Set up the Launcher
	config := &launcher.Config{
		AgentLoader: agent.NewSingleLoader(agentInstance),
	}

	l := full.NewLauncher()
	if err = l.Execute(ctx, config, os.Args[1:]); err != nil {
		log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
	}
}
