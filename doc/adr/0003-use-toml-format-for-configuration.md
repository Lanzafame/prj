# 3. Use toml format for configuration

Date: 28/02/2016

## Status

Accepted

## Context

To make `prj` useful to more people it should be able to be customized.

## Decision

Use .toml format for the configuration file as it is easy to understand/write/read.
Should make it an interface for easy swapping in of other formats.

## Consequences

Write config reader/writer interface. Only implement toml for now.
