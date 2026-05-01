# osctx - Interactive switch for Openstack clouds

## Problem

Imagine that you have multiple Openstack clouds and you want to switch between them easily. These clouds are described in your Openstack config file (clouds.yaml).

In some moment you need to get/create some entities in your **dev** cloud (for example) and after that you need to check your production environment virtual machines.

These actions needs to be performed via manual setting of `--os-cloud=cloud_name` or OS_CLOUD environment variable.

## Solution

`osctx` is a command line tool that allows you to switch between Openstack clouds easily. Project is inspired by [kubectx and kubens](https://github.com/ahmetb/kubectx) utitilies for Kubernetes.

## Requirements

- Works on Linux, MacOS and Windows
- Uses `fzf` (https://github.com/junegunn/fzf) for interactive selection
- If `fzf` is not installed, it will fallback to manual selection (by index in list of clouds)
- Sets `OS_CLOUD` environment variable for further usage of Openstack CLI
- Because there is no way to set environment variable for parent process in subprocess, `osctx` writes "export OS_CLOUD" to stdout. For full usage the user needs to source that output with the snippet below.
- All another output should go to stderr
- Uses `spf13/cobra` (https://github.com/spf13/cobra) for command line parsing

## Installation

For better experience you need to install `fzf` (https://github.com/junegunn/fzf).

1. Download binary file
2. Move it to your PATH
3. Add next line to your .bashrc or profile:
```bash
osctx() { eval "$(command osctx "$@")"; }
```

## Usage

#### Run interactive switch:

```sh
osctx
```

It will display you fuzzy finder inline window where you can find and interactive select your Openstack cloud from the list parsed from `cloud.yaml`.

#### Switch to previous cloud

```sh
osctx -
```

#### List all available clouds

```sh
osctx ls
```

#### Get current cloud

```sh
osctx current
```

#### Clear current cloud

```sh
osctx unset
```

## Libraries and external dependencies

- spf13/cobra https://github.com/spf13/cobra
- fzf (binary) https://github.com/junegunn/fzf