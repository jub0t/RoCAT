# RoCat - **Ro**blox **C**lothes **A**utomation **T**ool

RoCat is a revolutionary tool designed to simplify and automate the process of uploading cloth items to the Roblox catalog. With RoCat, users no longer need to manually upload each cloth item, saving them time and effort. RoCat is designed to be user-friendly and accessible to users of all levels of experience, making it easy for anyone to upload their own cloth creations to the Roblox catalog. The tool features an intuitive interface that guides users through the process, allowing them to quickly and easily upload their cloth items with just a few clicks.

```bash
NAME:
   RoCAT - Roblox clothing automation tool.

USAGE:
   RoCAT [global options] command [command options] [arguments...]

COMMANDS:
   info      Display information about the cli.
   whoami    Uses your cookie from the file and fetches account/bot info.
   download  Download classic clothing from roblox catalogue and save them for later upload
   start     Start uploading the stored clothing.
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

Download this tool for your OS from the [Releases](https://github.com/jareer12/RoCAT/releases/tag/stable) tab.

## Usage

When you run the CLI tool for the first time it will create a few new sub-directories, like `store` and `downloads` and `temp`.

### Download

Use this command to bulk download classic clothing from roblox catalogue. The cli tool will create a `downloads` folder in the same directory.

```sh
rocat download --amount 120 --type shirts
```

### Upload

Now you can upload the stored clothing to the website, the cli will keep track of the uploaded clothing, and will not re-upload clothing. Use the `--seo` flag to allow the cli to generate descriptions using an algorithm, better description can lead to more sales.

```
rocat start --groupId 7830839 --limit 10 --seo
```

## Build From Source

- For Linux Run the [Build Bash file](./scripts/build.sh)
- For Windows Run the [Build Batchfile](./scripts/build.cmd)

### Simple Build & Usage

```sh
## Init
go mod tidy

## Run the code directly
go run ./

## Build into exe(change name if you are not using windows)
go build -o ./out.exe
./out.exe
```

## Contribution

This tool is still under development and is not complete. Contributions & pull requests are appreciated.
