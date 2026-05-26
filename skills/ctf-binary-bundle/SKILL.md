---
name: ctf-binary-bundle
description: Bundle skill for CTF binary triage and solve routing. Use when the prompt mentions ELF, nc services, memory corruption, WASM reversing, native binaries, or uncertain binary challenge direction and Codex should coordinate the core CTF binary skills before deeper specialization.
metadata:
  version: "1.0.0"
  tags:
    - bundle
    - ctf
    - binary
    - pwn
    - reverse
    - wasm
  triggers:
    - elf
    - pwn
    - nc
    - binary challenge
    - reverse challenge
    - wasm challenge
  target_types:
    - ctf
    - binary
    - reverse
  bundle_of:
    - ctf-skills
    - elf-pwn-triage
    - wasm-reverse-triage
  recommended_tools:
    - strings
    - gdb
    - pwntools
    - ghidra
    - radare2
  role_hints:
    - CTF
    - Reverse engineer
    - Binary exploitation
  stages:
    - triage
    - verify
    - solve
  autoload_priority: 24
---

# CTF Binary Bundle

Use this bundle to route binary-heavy CTF prompts to the right next skill without overloading the context window.

## Workflow

1. Classify the sample as native pwn, reverse, or hybrid.
2. Load only the member skill that matches the actual blocker.
3. Keep evidence on architecture, mitigations, entrypoints, and remote constraints.

## Load Order

- Start with `ctf-skills` when the category is still ambiguous.
- Load `elf-pwn-triage` when the challenge is an ELF or an `nc`-backed native service.
- Load `wasm-reverse-triage` when the main blocker is browser or WASM logic recovery.

## Output

- Dominant subcategory
- Required member skills
- First-pass triage notes
- Next solve direction
