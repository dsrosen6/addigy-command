# addigy-command
This is a command line helper for various Addigy terminal commands - these are all commands that can already be run via commands in the Addigy admin panel, but if you are in a situation where it might make more sense to run it via terminal on a user's computer, this can help reduce the need to type out complex commands or filepaths. 

## Usage
Commands must be run as root, or with `sudo` - options are as follows.

### Menu
Simply running `addigy` will pull up an interactive menu that lets you select from one of the following commands.

### Run
`addigy run`: Run the Addigy policy. Without arguments, this will start the policy with no progress spinner.  
`-s, --spinner`: Run the Addigy policy with a progress spinner, which stops when the policy run is complete.  
`-v, --verbose`: Run the Addigy policy with full verbose output.

### Reset
`addigy reset`: Reset the progress of all Addigy policy items. Use as a "flush" if you need to reset the progress of all policy items, but time is not of the essence.

### Full Reset
`addigy full-reset`: Reset the Addigy policy progress (reset), and run the policy (run). Policy will run without a spinner.  
`-s, --spinner`: Policy will run with a progress spinner (`run -f`).  
`-v, --verbose`: Policy will run with full verbose output (`run -v`).  

### Help
`-h, --help, help` Show the help message (includes all of the above)