<!-- distilled: 2026-05-12T08:06:24.129Z | source: AGENTS.md -->

## Purpose
This document outlines the operational rules and workflow constraints for a coding agent called Pi. It establishes guidelines to ensure concise, file-based reasoning in project collaborations.

## Key Exports
- **No direct exports**: This is a metadata file specifying agent behavior, not code to be executed.

## Dependencies
- Depends on external files: 
  - **/opt/homebrew/lib/node_modules/@mariozechner/pi-coding-agent/README.md**: Main documentation file, referenced for initial instructions.
  - **Project-specific files under /Users/davidsimmons/Languages/elixir/phoenix/guess/.think/****: Used for storing task state, plans, step-by-step reasoning logs (`_plan.md`, `step-NNN.md`), and summaries.

## Key Logic
- **Core Principle**: All reasoning must occur by writing or reading files, not mentally. Use `.think/` directory to manage state.
- **Strict Workflow**: 
  - Every turn must start by reading `_state.md`.
  - For a new task, first create `_state.md` at session start.
- **HARD RULES**:  
  - Do NOT hold entire codebase context — read only one file at a time.
  - Always update `_state.md` — it serves as the task lifeline between turns.

## Integration Points
- **File-based Workflow**: Reads and writes to `.think/` files only. Relies on cross-references in PI documentation for core processes.
- **PI's Knowledge**: If relevant, refer to `~/.pi/knowledge/` during initial KNOWLEDGE SCAN.

## Summary
Pi's operational logic focuses on task separation: first define the state via `_state.md`, then handle analysis or code edits by writing step files. This AI agent remembers only via the filesystem to maintain traceability and prevent context loss, enforcing a short-response format for efficient turns.