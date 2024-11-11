# FileBackupToolCLI

FileBackupToolCLI is a command-line tool developed in Go. This tool allows you to back up a source file to a target path at specified intervals. The application provides a user-friendly interface using the `cobra` package.

## Features

- Copy a source file to a target path at specified intervals.
- Define source and target paths, as well as the copy frequency.
- Schedule the copy operation in seconds.

## Installation and Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/barisaydogdu/FileBackupToolCLI.git
   ```

2. Navigate to the project directory:
   ```bash
   cd FileBackupToolCLI
   ```

3. Install the required dependencies:
   ```bash
   go mod tidy
   ```

4. Build the project:
   ```bash
   go build -o filecopy
   ```

5. Run the tool:
   ```bash
   ./filecopy
   ```

## Usage

Use the `backupfile` command to specify the source file, target file, and copy frequency:

```bash
./filecopy backupfile --sourcefile /path/to/source/file --targetfile /path/to/target/file --period 60
```

### Command Line Options

- `--sourcefile`: Path to the source file you want to copy.
- `--targetfile`: Path to the target location.
- `--period`: Copy frequency in seconds.

### Example

To copy `/home/user/source.txt` to `/home/user/backup.txt` every 120 seconds:

```bash
./filecopy backupfile --sourcefile /home/user/source.txt --targetfile /home/user/backup.txt --period 120
```

## Technologies Used

- **Go**: The programming language used for the project.
- **Cobra**: Library used for building the command-line interface.

## Contributing

If you would like to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b feature-name
   ```
3. Make your changes and commit them:
   ```bash
   git commit -m 'Added new feature'
   ```
4. Push your branch to the remote repository:
   ```bash
   git push origin feature-name
   ```
5. Create a Pull Request.
