# Hexedit

> THis project is largely just for play & practise and is in a development state :-)

hexedit is a terminal application which helps read, write and update hex data.  This is largely just for play & practise

It is heavily inspired by the fanastic VSCode extension https://marketplace.visualstudio.com/items?itemName=ms-vscode.hexeditor

## Usage

```
hexedit -f <filepath> [-o] [-h]
```

### Options:
- `-f <filepath>`: Path to the input file (required)
- `-o`: Output the contents and exit without entering interactive mode
- `-h`: Display help information and exit

### Examples:

```bash
# Open a file in interactive mode
hexedit -f myfile.bin

# Output file contents and exit
hexedit -f myfile.bin -o

# Display help information
hexedit -h
```

