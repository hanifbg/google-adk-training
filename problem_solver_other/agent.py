"""
Problem-solving agent with built-in planning capabilities.

Demonstrates ADK's BuiltInPlanner with ThinkingConfig.
"""

import os

# Manual environment variable loading from .env if it exists
def load_dotenv(path=".env"):
    if os.path.exists(path):
        with open(path, "r") as f:
            for line in f:
                line = line.strip()
                if not line or line.startswith("#"):
                    continue
                if "=" in line:
                    key, value = line.split("=", 1)
                    # Remove optional quotes
                    value = value.strip().strip("'").strip('"')
                    os.environ[key.strip()] = value

# Load the .env file
load_dotenv()

from google.adk.agents import LlmAgent
from google.adk.planners import BuiltInPlanner
from google.genai import types
from google.adk.planners import PlanReActPlanner
from google.adk.models.lite_llm import LiteLlm

# Use 'openai/' prefix to tell LiteLLM to use the OpenAI-compatible provider
# for your custom endpoint.
remote_llm = LiteLlm(
    model="openai/seed-2-0-lite-free",
    api_base=os.getenv("BASE_API"),
    api_key=os.getenv("API_KEY")
)

# Planning-enabled agent for complex problem solving
root_agent = LlmAgent(
    model=remote_llm,
    name="strategic_problem_solver",
    description="Solves complex problems using multi-step reasoning and planning",
    instruction="""You are a Strategic Problem Solver.

Your approach to complex problems:
1. **Understand** - Break down the problem into components
2. **Analyze** - Consider multiple approaches and trade-offs
3. **Plan** - Develop a step-by-step solution strategy
4. **Execute** - Provide clear, actionable recommendations

For complex problems:
- Think through implications and edge cases
- Consider short-term vs long-term consequences
- Identify potential risks and mitigation strategies
- Provide reasoning for your recommendations

Be thorough, analytical, and systematic in your approach.""",
    planner=PlanReActPlanner()
)