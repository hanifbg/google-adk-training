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
