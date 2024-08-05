package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/dsrosen6/addigy-command/pkg/addigy"
)

var usage = `
Usage: addigy [command]

Commands:
  run           Run the Addigy policy. Without arguments, this will start the policy with no progress spinner.
  				Arguments:
				-s      Run the Addigy policy with a progress spinner, which stops when the policy run is complete.
				-v	    Run the Addigy policy with full verbose output.
				 
  reset       Reset the progress of all Addigy policy items. Use as a "flush" if you need to reset the progress of all policy items, but time is not of the essence.

  full-reset  Reset the Addigy policy progress (reset), and run the policy (run). Policy will run without a spinner.
				-s	  Policy will run with a progress spinner (run -f).
				-v	  Policy will run with full verbose output (run -v).

  help        Show this help message.
`

func Run() error {
	// if there are any args on top of the binary name, process them
	if len(os.Args) > 1 {
		processArgs()
		return nil
	} else {
		return runCommandMenu()
	}
}

func processArgs() error {
	switch os.Args[1] {

	case "run":
		if len(os.Args) > 2 {
			switch os.Args[2] {
			case "-s":
				return addigy.PolicierRunWithSpinner()
			case "-v":
				return addigy.PolicierRunVerbose()
			default:
				return errors.New("invalid argument - use 'addigy help' for usage")
			}
		} else {
			addigy.PolicierRun()
		}

	case "reset":
		return addigy.ResetPolicyProgress()

	case "full-reset":
		if len(os.Args) > 2 {
			switch os.Args[2] {
			case "-s":
				return FullResetSpinner()
			case "-v":
				return FullResetVerbose()
			default:
				return errors.New("invalid argument - use 'addigy help' for usage")
			}
		} else {
			return FullReset()
		}

	case "help":
		fmt.Println(usage)

	default:
		fmt.Println("Invalid command. Use 'addigy help' for usage.")
	}

	return nil
}

func runCommandMenu() error {
	var selection string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Value(&selection).
				Height(8).
				Title("Command Description").
				Options(
					huh.NewOption("Run Policy", "run"),
					huh.NewOption("Run Policy (Spinner)", "run-s"),
					huh.NewOption("Run Policy (Verbose)", "run-v"),
					huh.NewOption("Reset Policy Progress", "reset"),
					huh.NewOption("Reset Progress + Run Policy", "reset-run"),
					huh.NewOption("Reset Progress + Run Policy (Spinner)", "reset-run-s"),
					huh.NewOption("Reset Progress + Run Policy (Verbose)", "reset-run-v"),
				).
				DescriptionFunc(func() string {
					switch selection {
					case "run":
						return "Start the Addigy policy run.\nQuick jumpstart - no spinner.\n"
					case "run-s":
						return "Start the Addigy policy run.\nRuns with a spinner, which stops when the policy is done.\n"
					case "run-v":
						return "Start the Addigy policy run.\nRuns with full verbose output.\n"
					case "reset":
						return "Reset the progress of all Addigy policy items.\nUse this if you need to flush the progress of all policy items, but time is not of the essence.\n"
					case "reset-run":
						return "Reset the Addigy policy progress, and run the policy.\nPolicy will run without a spinner.\n"
					case "reset-run-s":
						return "Reset the Addigy policy progress, and run the policy.\nPolicy will run with a spinner.\n"
					case "reset-run-v":
						return "Reset the Addigy policy progress, and run the policy.\nPolicy will run with full verbose output.\n"
					default:
						return "Select an option.\n"
					}
				}, &selection),
		),
	)

	if err := form.Run(); err != nil {
		return err
	}

	switch selection {
	case "run":
		return addigy.PolicierRun()
	case "run-s":
		return addigy.PolicierRunWithSpinner()
	case "run-v":
		return addigy.PolicierRunVerbose()
	case "reset":
		return addigy.ResetPolicyProgress()
	case "reset-run":
		return FullReset()
	case "reset-run-s":
		return FullResetSpinner()
	case "reset-run-v":
		return FullResetVerbose()
	default:
		return nil
	}
}

func FullReset() error {
	err := addigy.ResetPolicyProgress()
	if err != nil {
		return err
	}
	return addigy.PolicierRun()
}

func FullResetSpinner() error {
	err := addigy.ResetPolicyProgress()
	if err != nil {
		return err
	}
	return addigy.PolicierRunWithSpinner()
}

func FullResetVerbose() error {
	err := addigy.ResetPolicyProgress()
	if err != nil {
		return err
	}
	return addigy.PolicierRunVerbose()
}
