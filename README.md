# addigy-command
This is a command line helper for various Addigy terminal commands - these are all commands that can already be run via commands in the Addigy admin panel, but if you are in a situation where it might make more sense to run it via terminal on a user's computer, this can help reduce the need to type out complex commands or filepaths. 

## Usage
Commands must be run as root, or with `sudo` - options are as follows.

### Menu
Simply running `addigy` will pull up an interactive menu that lets you select from one of the following commands.

### Run
`addigy run`: Run the Addigy policy (starts with no further output)
`-s, --spinner`: Run the Addigy policy with a progress spinner
`-v, --verbose`: Run the Addigy policy with full verbose output

### Reset
`addigy reset`: Reset the progress of all Addigy policy items

### Full Reset
`addigy full-reset`: Runs "reset" and "run" in sequence (use as a full flush)	
`-s, --spinner`: Policy will run with a progress spinner (`run -s`)
`-v, --verbose`: Policy will run with full verbose output (`run -v`)

### Help
`-h, --help, help` Show the help message (includes all of the above)