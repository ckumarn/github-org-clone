package main

import (
	"os"
	"github.com/urfave/cli"
	"github.com/steinfletcher/github-team-clone/github"
	"github.com/steinfletcher/github-team-clone/teamclone"
	"github.com/apex/log"
)

func main() {
	app := cli.NewApp()
	app.Author = "Stein Fletcher"
	app.Name = "github-team-clone"
	app.Usage = "clone github team repos"
	app.UsageText = "github-team-clone -o MyOrg -t MyTeam"
	app.Version = "0.0.1"
	app.EnableBashCompletion = true
	app.Description = "A simple cli to clone all the repos managed by a github team"

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "org, o",
			Usage: "github organisation",
		},
		cli.StringFlag{
			Name: "team, t",
			Usage: "github team",
		},
		cli.StringFlag{
			Name: "username, u",
			Usage: "github username",
			EnvVar: "GITHUB_USER,GITHUB_USERNAME",
		},
		cli.StringFlag{
			Name: "token, k",
			Usage: "github personal access token",
			EnvVar: "GITHUB_TOKEN,GITHUB_API_KEY,GITHUB_PERSONAL_ACCESS_TOKEN",
		},
		cli.StringFlag{
			Name: "dir, d",
			Usage: "directory to clone into",
			Value: "src",
		},
	}

	app.Action = func(c *cli.Context) error {
		username := c.String("username")
		token := c.String("token")
		team := c.String("team")
		org := c.String("org")
		dir := c.String("dir")

		if len(username) == 0 {
			die("env var GITHUB_USERNAME or flag -u must be set", c)
		}

		if len(token) == 0 {
			die("env var GITHUB_TOKEN or flag -k must be set", c)
		}

		if len(team) == 0 {
			die("github team (-t) not set", c)
		}

		if len(org) == 0 {
			die("github organisation (-o) not set", c)
		}

		githubCli := github.NewGithub(username, token)
		cloner := teamclone.NewCloner(githubCli, dir)

		err := cloner.CloneTeamRepos(org, team)
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		return nil
	}

	app.Run(os.Args)
}

func die(msg string, c *cli.Context) {
	cli.ShowAppHelp(c)
	log.Fatal(msg)
}