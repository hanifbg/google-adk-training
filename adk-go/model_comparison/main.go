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

func main() {
	ctx := context.Background()

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	clientConfig := &genai.ClientConfig{
		Backend:  genai.BackendVertexAI,
		Project:  os.Getenv("GOOGLE_CLOUD_PROJECT"),
		Location: os.Getenv("GOOGLE_CLOUD_LOCATION"),
	}

	// Agent 1: Optimized for Factual Data Extraction
	factualModel, err := gemini.NewModel(ctx, "gemini-2.5-flash", clientConfig)
	if err != nil {
		log.Fatalf("Failed to create factual model: %v", err)
	}

	factualAgentConfig := llmagent.Config{
		Name:        "data_extractor",
		Model:       factualModel,
		Description: "Extracts factual information with high consistency",
		Instruction: `You are a precise data extractor.

Extract facts exactly as stated. Do not:
- Add information not present in the input
- Make assumptions or inferences
- Use creative language

Be accurate, concise, and deterministic.`,
		GenerateContentConfig: &genai.GenerateContentConfig{
			Temperature:     ptr(float32(0.1)),
			MaxOutputTokens: 500,
			TopP:            ptr(float32(0.8)),
			TopK:            ptr(float32(10)),
			SafetySettings: []*genai.SafetySetting{
				{
					Category:  genai.HarmCategoryDangerousContent,
					Threshold: genai.HarmBlockThresholdBlockLowAndAbove,
				},
			},
		},
	}

	factualAgent, err := llmagent.New(factualAgentConfig)
	if err != nil {
		log.Fatalf("Failed to create factual agent: %v", err)
	}

	// Agent 2: Optimized for Creative Brainstorming
	creativeModel, err := gemini.NewModel(ctx, "gemini-2.5-pro", clientConfig)
	if err != nil {
		log.Fatalf("Failed to create creative model: %v", err)
	}

	creativeAgentConfig := llmagent.Config{
		Name:        "creative_brainstormer",
		Model:       creativeModel,
		Description: "Generates creative ideas and explores possibilities",
		Instruction: `You are a creative brainstorming partner.

Generate innovative, diverse, and imaginative ideas. Feel free to:
- Think outside the box
- Combine unexpected concepts
- Explore unconventional approaches

Be creative, varied, and thought-provoking.`,
		GenerateContentConfig: &genai.GenerateContentConfig{
			Temperature:     ptr(float32(0.9)),
			MaxOutputTokens: 2000,
			TopP:            ptr(float32(0.95)),
			TopK:            ptr(float32(40)),
			SafetySettings: []*genai.SafetySetting{
				{
					Category:  genai.HarmCategoryDangerousContent,
					Threshold: genai.HarmBlockThresholdBlockMediumAndAbove,
				},
			},
		},
	}

	creativeAgent, err := llmagent.New(creativeAgentConfig)
	if err != nil {
		log.Fatalf("Failed to create creative agent: %v", err)
	}

	loader, err := agent.NewMultiLoader(factualAgent, creativeAgent)
	if err != nil {
		log.Fatalf("Failed to create agent loader: %v", err)
	}

	config := &launcher.Config{
		AgentLoader: loader,
	}

	l := full.NewLauncher()
	if err = l.Execute(ctx, config, os.Args[1:]); err != nil {
		log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
	}
}

func ptr[T any](v T) *T {
	return &v
}
