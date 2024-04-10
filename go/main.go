// homechart is a household management platform.
package main

import (
	"os"

	"github.com/candiddev/homechart/go/config"
	"github.com/candiddev/shared/go/cli"
)

//nolint:gochecknoglobals
var (
	appCloudPublicKey = "ed25519public:MCowBQYDK2VwAyEADL5OxQve4AvYy7L2S+ypqD0/T8t9IIT/bQFkXNQCo9I="
)

func main() {
	if err := (&cli.App[*config.Config]{
		Commands: map[string]cli.Command[*config.Config]{
			"generate-cloud": {
				Run: generateCloud,
			},
			"gen-keys": {
				Run: cli.GenKeys[*config.Config]().Run,
			},
			"generate-vapid": {
				Run:   generateVAPID,
				Usage: "Generate a VAPID public and private key for Web Push",
			},
			"run": {
				Run:   run,
				Usage: "Start Homechart API server",
			},
			"seed": {
				ArgumentsRequired: []string{
					"output path",
				},
				Run:   seed,
				Usage: "Seed the database with mock data and save the output as JSON to path",
			},
			"tasks-day": {
				Run:   tasksRun,
				Usage: "Manually run tasks that occur every day",
			},
			"tasks-hour": {
				Run:   tasksRun,
				Usage: "Manually run tasks that occur every hour",
			},
			"tasks-minute": {
				Run:   tasksRun,
				Usage: "Manually run tasks that occur every minute",
			},
			"tasks-minute-five": {
				Run:   tasksRun,
				Usage: "Manually run tasks that occur every five minutes",
			},
		},
		Config:      config.Default(),
		Description: "Homechart runs your household.",
		HideConfigFields: []string{
			"app.cloudEndpoint",
			"app.keepExpiredAuthHouseholdsDays",
			"app.trialDays",
			"apple",
			"android",
			"fcm",
			"google",
			"oauth",
			"oidc",
			"paddle",
			"tracing",
		},
		Name:        "Homechart",
		PricingLink: "https://homechart.app/pricing",
	}).Run(); err != nil {
		os.Exit(1)
	}
}
