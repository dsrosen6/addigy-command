package cli

import (
	"os"

	"github.com/charmbracelet/huh"
	"github.com/dsrosen6/addigy-command/pkg/addigy"
)

var usage = `
Usage: addigy [command]

Commands:
  run         Run the Addigy policy, and just let you know if it started.
  run -f      Run the Addigy policy with a spinner which stops when the policy run is complete.
  reset       Reset the Addigy policy progress.
  flush       Reset the Addigy policy progress, and run the policy.
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

func processArgs() {
	switch os.Args[1] {
	case "run":
		if len(os.Args) > 2 && os.Args[2] == "-f" {
			addigy.PolicierRunWithSpinner()
		} else {
			addigy.PolicierRun()
		}
	case "reset":
		addigy.ResetPolicyProgress()
	case "flush":
		if len(os.Args) > 2 && os.Args[2] == "-f" {
			addigy.ResetPolicyProgress()
			addigy.PolicierRunWithSpinner()
		} else {
			addigy.ResetPolicyProgress()
			addigy.PolicierRun()
		}
	case "help":
		println(usage)
	default:
		runCommandMenu()
	}
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
					huh.NewOption("Run Policy (Full)", "run-full"),
					huh.NewOption("Reset Policy Progress", "reset"),
					huh.NewOption("Reset Progress + Run Policy", "reset-run"),
				).
				DescriptionFunc(func() string {
					switch selection {
					case "run":
						return "Run the Addigy policy, and just let you know if it started.\nGood if you just need a quick jumpstart, but don't need to know when it's done.\n"
					case "run-full":
						return "Run the Addigy policy with a spinner which stops when the policy run is complete.\nIf you need confirmation that the policy is done running, use this.\n"
					case "reset":
						return "Reset the Addigy policy progress.\nUse this if you need to flush the progress of all policy items, but time is not of the essence.\n"
					case "reset-run":
						return "Reset the Addigy policy progress, and run the policy.\nUse this if you need to flush the progress of all policy items, and need to run the policy immediately.\n"
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
	case "run-full":
		return addigy.PolicierRunWithSpinner()
	case "reset":
		return addigy.ResetPolicyProgress()
	case "reset-run":
		if err := addigy.ResetPolicyProgress(); err != nil {
			return err
		}
		return addigy.PolicierRun()
	default:
		return nil
	}
}
