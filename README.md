# 🤖 google-adk-training

This repository contains training materials and exercises for building AI agents using the **Google Agent Development Kit (ADK)**. 
It follows the official training path [Build Agents with the Agent Development Kit (ADK)](https://www.skills.google/paths/3545).

## 📂 Project Structure

This workspace includes two sub-agent projects demonstrating different ways to define agents:

### 1. 🧮 Math Tutor Agent (`my_first_agent`)
- **Type**: Python-defined agent.
- **File**: `my_first_agent/agent.py`
- **Goal**: A patient math tutor that helps students learn algebra by guiding them through problem-solving steps.
- **Model**: `gemini-2.5-flash`

### 2. 🗣️ Assistant Agent (`my_agent`)
- **Type**: YAML-defined agent.
- **File**: `my_agent/root_agent.yaml`
- **Goal**: A helpful assistant for various tasks.
- **Model**: `gemini-2.5-flash`

### 3. 🧠 Strategic Problem Solver (`problem_solver_gemini`)
- **Type**: Python-defined agent with planning.
- **File**: `problem_solver_gemini/agent.py`
- **Goal**: Solves complex problems using a step-by-step reasoning strategy (`BuiltInPlanner`).
- **Model**: `gemini-2.5-flash`

### 4. 🌐 Custom Provider Agent (`problem_solver_other`)
- **Type**: Python-defined agent with planning.
- **File**: `problem_solver_other/agent.py`
- **Goal**: Demonstrates calling non-Gemini models via custom OpenAI-compatible endpoints (`LiteLlm`).
- **Model**: `openai/seed-2-0-lite-free` (via Sumopod)

---

## 🛠️ Setup Instructions

### 1. Create a Virtual Environment
It is recommended to use a virtual environment for each project:
```bash
python3 -m venv .venv
source .venv/bin/activate
```

### 2. Install Dependencies
You'll need the `google-adk` package to run your agents:
```bash
pip install google-adk
```

### 3. Configure API Keys
Copy the `.env.template` (if provided) or create a `.env` file in each agent's directory with your Google Cloud or Gemini API credentials.
```bash
GOOGLE_API_KEY=YOUR_API_KEY_HERE
```

#### Custom Providers (e.g., Sumopod, vLLM, Groq)
For agents using `LiteLlm` (like `problem_solver_other`), configure your `.env` as follows:
```bash
BASE_API=https://ai.example.com/v1
API_KEY=YOUR_CUSTOM_KEY
```
> [!TIP]
> When using a custom `BASE_API`, prepend the provider prefix to your model name (e.g., `model="openai/my-model-name"`) to ensure LiteLLM routes the request correctly.

#### Common LiteLLM Providers
| Provider | Model Prefix | Typical `BASE_API` |
| :--- | :--- | :--- |
| **OpenAI Compatible** | `openai/` | `https://api.example.com/v1` |
| **Anthropic** | `anthropic/` | `https://api.anthropic.com` |
| **Ollama** | `ollama/` | `http://localhost:11434` |
| **Anyscale** | `anyscale/` | `https://api.endpoints.anyscale.com/v1` |
| **DeepSeek** | `deepseek/` | `https://api.deepseek.com` |
| **Groq** | `groq/` | `https://api.groq.com/openai/v1` |
| **Mistral** | `mistral/` | `https://api.mistral.ai/v1` |
| **Perplexity** | `perplexity/` | `https://api.perplexity.ai` |

---

## 🚀 Running Your Agents

You can interact with your agents using the ADK CLI:

### Using Python Agent
```bash
adk run my_first_agent
```

### Using YAML Agent
```bash
adk run my_agent
```

---

## 📚 Resources
- [ADK Official Documentation](https://github.com/google/adk-python)
- [Gemini Developer Center](https://ai.google.dev/)
- [Skills Boost Path: Build Agents with ADK](https://www.skills.google/paths/3545)
