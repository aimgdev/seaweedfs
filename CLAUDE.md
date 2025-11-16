# SeaweedFS - Claude Code Instructions

## Project Overview
SeaweedFS is a simple and highly scalable distributed file system designed to:
1. Store billions of files
2. Serve files fast

This is a fork maintained at `git@github.com:aimgdev/seaweedfs.git` with custom enhancements.

## Current Branch
- **Branch**: `public-http-enhancements`
- **Current Version**: jet r4
- **Latest Commit**: Add path-based access control for readonly port with per-path file and listing permissions

## Version Naming Convention
- Use incrementing revision numbers on the codename (e.g., jet r1, jet r2, jet r3, etc.)
- **IMPORTANT**: Increment the version number EVERY TIME before committing/pushing changes
- This allows easy tracking of what version is deployed
- Version numbers only go UP, never down or reused
- Update the version in two places:
  1. `/weed/util/version/constants.go` - Update the `CUSTOM_BUILD` variable
  2. This file (CLAUDE.md) - Update the "Current Version" field above

## Recent Enhancements
- Path-based access control for readonly port (per-path file/listing permissions) - jet r4
- Security improvements for directory listings (404 vs 403) - jet r3
- HTTP throughput optimization with larger buffers for high-speed networks - jet r2
- Readonly port directory listing control - jet r1
- Context-aware S3 action resolution

## Development Notes
- Main binary: `weed` (compiled from `/weed` directory)
- Language: Go
- Architecture: Distributed file system with master servers, volume servers, and filers
- S3 API compatibility included

## Important Files
- `README.md` - Main project documentation
- `DESIGN.md` - Architecture and design documentation
- `Makefile` - Build configuration
- `weed/` - Main source directory

## Build & Development
- Standard Go project structure
- Use `make` for building
- Docker support available in `docker/` directory
- Kubernetes configs in `k8s/` directory

## Upstream
- Upstream repo: `https://github.com/seaweedfs/seaweedfs.git`
- Keep track of upstream changes and selectively merge relevant updates
