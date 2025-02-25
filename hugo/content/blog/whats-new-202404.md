---
author: Mike
date: 2024-04-01
description: Release notes for Homechart v2024.04.
tags:
  - release
title: "What's New in Homechart: v2024.04"
type: blog
---

{{< homechart-release version="2024.04" >}}

## Enhancements

- Added SBOMs to Homechart releases.  See [Self-Hosted Homechart]({{< ref "/docs/guides/get-homechart/self-hosted#sbom" >}}) for more information.
- Updated Go to 1.22.1.

## Removals

- Removed support for old environment variable format, such as `HOMECHART_CLI_DEBUG`, will no longer work.  Please convert your environment variables to the new format, `HOMECHART_cli_logLevel`.  Visit [the config docs]({{< ref "/docs/references/config" >}}) for more information.
