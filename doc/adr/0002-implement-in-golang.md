# 2. Implement in Golang

Date: 18/02/2016

## Status

Accepted

## Context

The tool needs to create new files and git repos and apply small edits to the Status section of existing files.

## Decision

The tool is implemented in Golang as it has a libgit2 wrapper, a license generator and easy file creation.

## Consequences

This means re-writing [ADR-Tools][https://github.com/npryce/adr-tools] in Golang so that Windows is also supported.
