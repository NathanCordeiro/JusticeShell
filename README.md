# JusticeShell
JusticeShell is a fun, interactive command-line tool that lets you create dialogue between DC Comics characters.

## Installation, Building & Running

```bash
# Clone the repository
git clone https://github.com/NathanCordeiro/JusticeShell
cd JusticeShell
```

### Windows

1. Open Command Prompt or PowerShell
2. Navigate to the extracted/cloned directory:
```cmd
cd path\to\JusticeShell
```
3. Build the executable:
```cmd
go build -o JusticeShell.exe
```
4. Run the program:
```cmd
# From the same directory
JusticeShell.exe --speaker batman --message "I am vengeance!"

# To have characters talk to each other
JusticeShell.exe --speaker joker --message "Why so serious?" --respond batman
```

### macOS/Linux

1. Open Terminal
2. Navigate to the extracted/cloned directory:
```bash
cd path/to/JusticeShell
```
3. Build the executable:
```bash
go build -o JusticeShell
```
4. Run the program:
```bash
# From the same directory
./JusticeShell --speaker batman --message "I am vengeance!"

# To have characters talk to each other
./JusticeShell --speaker joker --message "Why so serious?" --respond batman
```


### Available Characters

- **batman**
- **joker**
- **superman**
- **wonder_woman**

### Command Line Flags

- `--speaker`: Choose the speaking character (default: "batman")
- `--message`: Set the character's dialogue (default: "I am vengeance!")
- `--respond`: Optional: Have another character respond

## Examples

Batman monologue:

```bash
./JusticeShell --speaker batman --message "The night is darkest before the dawn."
```

Joker taunting Batman:

```bash
./JusticeShell --speaker joker --message "Let's put a smile on that face!" --respond batman
```

Superman inspiring hope:

```bash
./JusticeShell --speaker superman --message "Up, up and away!"
```

## Acknowledgments

- Inspired by the classic `cowsay` Unix program
- ASCII art adapted from various sources
