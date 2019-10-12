# Labeler

Labeler is a CLI tool to sync labels for a GitHub repository with templates.

## Usage

```
NAME:
   labeler - A CLI tool to sync labels for a GitHub repository with templates

USAGE:
    [global options] command [command options] [arguments...]

VERSION:
   0.2.0

COMMANDS:
   save     Save labels of target repository to a template file
   sync     Sync labels from a template file to target repository
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --token value  GitHub personal access token [$LABELER_TOKEN]
   --help, -h     show help
   --version, -v  print the version
```

## Example

Save labels from a repository:
```sh
$ export LABELER_TOKEN={your GitHub token}
$ labeler save --owner=unknwon --repo=labeler --to unknwon_labeler.json
```

Sync labels to a repository:
```sh
$ export LABELER_TOKEN={your GitHub token}
$ labeler sync --owner=unknwon --repo=labeler --from unknwon_labeler.json
```

## License

This project is under MIT License. See the [LICENSE](LICENSE) file for the full license text.
