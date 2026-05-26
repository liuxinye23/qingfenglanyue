---
name: forensics-triage-bundle
description: Bundle skill for digital forensics and artifact-first triage. Use when requests involve PCAPs, stego, suspicious images, metadata, extracted files, or mixed forensic artifacts and Codex should coordinate the initial triage skills before deeper analysis.
metadata:
  version: "1.0.0"
  tags:
    - bundle
    - forensics
    - pcap
    - stego
    - artifact
    - ctf
  triggers:
    - pcap
    - packet capture
    - stego
    - suspicious image
    - metadata
    - forensic artifact
  target_types:
    - forensics
    - ctf
  bundle_of:
    - ctf-skills
    - pcap-triage
    - stego-triage
  recommended_tools:
    - binwalk
    - exiftool
    - foremost
    - zsteg
    - volatility3
  role_hints:
    - Forensics
    - Incident response
    - CTF
  stages:
    - triage
    - verify
    - solve
  autoload_priority: 23
---

# Forensics Triage Bundle

Use this bundle when the input is an artifact collection and the first job is to decide where the signal is.

## Workflow

1. Preserve artifact boundaries and classify by type.
2. Load only the member skill that matches the dominant artifact.
3. Record extracted objects, timestamps, and suspicious pivots for handoff.

## Load Order

- Start with `ctf-skills` when the category is still mixed or uncertain.
- Load `pcap-triage` for traffic captures and exported network evidence.
- Load `stego-triage` for images, embedded objects, and hidden-content workflows.

## Output

- Artifact classification
- Member skills loaded
- Highest-signal pivots
- Next analysis direction
