[![asciicast](https://asciinema.org/a/ds04hF3a8sDnaHjbMQujrf85g.svg)](https://asciinema.org/a/ds04hF3a8sDnaHjbMQujrf85g)

# tp

⚡️ Teleport to your projects and run commands in an instant.

## Introduction

tp is a simple command line tool aimed to help you streamline your development experience. Instead of having to run multiple commands — including changing your directory and running additional commands —, tp aims to consolidate all these commands into one simple one; possibly saving seconds of your time and more.

## Motivation

Having had experience with several projects involving Node.js and npm, I realised that I had to enter quite a few commands before being able to get the project running. For instance, I had to first change my working directory to the project's directory then run an npm script.

I realised that it could be better by combining the two commands together, but even so it might take a while for me to type the entire chain of commands out every time.

I settled on creating tp to help me with this issue. I decided to use Go knowing its fairly fast execution, and used this project as a primer into the world of Go as well.

## Installation

There are two methods to installing tp:

1. [through Homebrew](#installing-through-homebrew); and
2. by [manually installing](#installing-manually).

### Installing through Homebrew

To install through Homebrew, access the tap at `arashnrim/brew`:

```
brew install arashnrim/tap/tp
```

### Installing manually

To get started, visit [the Releases tab](https://github.com/arashnrim/tp/releases) and download the appropriate build for your operating system.

Next, just to confirm, update your permissions to ensure that the executable file can be run.

```
chmod +x path/to/tp
```

Next, move the executable file to a location that is added to your terminal's `$PATH` environment variable. An instance could be `/usr/local/bin/`, where executable binaries are usually located.

```
mv path/to/tp $path/added/directory
```

## Usage

To see the list of commands available, run the `help` command:

```
tp help
```

There are four commands that you may run: `add`, `remove`, `list`, and `to`.

The first command, `add`, allows you to set up a new location for tp:

```
tp <name> <location>
```

The second, `remove`, deletes an existing location from tp. This will not delete the actual location:

```
tp remove <name>
```

The third, `list`, shows all the available locations that you have set up:

```
tp list
```

The last, `to`, teleports to the location and runs commands there. This will not change your current working directory to the folder, though:

```
tp to <name>
```

## Licence

This repository is made open-source with the [MIT License](https://github.com/arashnrim/tp/blob/main/LICENSE.md), meaning that you are allowed to modify and distribute the source code as well as use it for private and commercial use provided that the licence and copyright notices are retained. For more information, visit the link above to learn what is permitted by the licence.
