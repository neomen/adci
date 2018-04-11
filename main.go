package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"time"

	"github.com/tucnak/climax"
)

const mysqlRootPass string = "root"

func randomString(n int) string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func runcmd(cmd string) {
	fmt.Println(cmd)
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
}

func main() {
	cli := climax.New("ADCI")
	cli.Brief = "adci clients server managment app"
	cli.Version = "beta 2.0"

	addCmd := climax.Command{
		Name:  "add",
		Brief: "merges the strings given",
		Usage: `[-s=] "a few" distinct strings`,
		Help:  `Lorem ipsum dolor sit amet amet sit todor...`,

		Flags: []climax.Flag{
			{
				Name:     "user",
				Short:    "u",
				Usage:    `--user="userName"`,
				Help:     `Put some username string .`,
				Variable: true,
			},
			{
				Name:     "host",
				Short:    "h",
				Usage:    `--host="hostName"`,
				Help:     `Put some host name string .`,
				Variable: true,
			},
			// {
			// 	Name:     "database",
			// 	Short:    "db",
			// 	Usage:    `--database`,
			// 	Help:     `this flag create `,
			// 	Variable: true,
			// },
		},

		Examples: []climax.Example{
			{
				Usecase:     `-u userName`,
				Description: `this flag adds MySQL user and Linux system user`,
			},
			{
				Usecase:     `-h hostName`,
				Description: `this flag adds Nginx virtual host `,
			},
		},

		Handle: func(ctx climax.Context) int {
			// var separator string
			if user, ok := ctx.Get("user"); ok {
				// separator = sep
				addUser(user)
				fmt.Println(user)
			}
			if host, ok := ctx.Get("host"); ok {
				// separator = sep
				addHost(host, "d7")
				//fmt.Println(user)
			}
			//fmt.Println(strings.Join(ctx.Args, separator))

			return 0
		},
	}
	deleteCmd := climax.Command{
		Name:  "delete",
		Brief: "merges the strings given",
		Usage: `[-s=] "a few" distinct strings`,
		Help:  `Lorem ipsum dolor sit amet amet sit todor...`,

		Flags: []climax.Flag{
			{
				Name:     "user",
				Short:    "u",
				Usage:    `--user="userName"`,
				Help:     `Put some username string .`,
				Variable: true,
			},
		},

		Examples: []climax.Example{
			{
				Usecase:     `-u userName`,
				Description: `Results in "google.com"`,
			},
		},

		Handle: func(ctx climax.Context) int {
			// var separator string
			if user, ok := ctx.Get("user"); ok {
				// separator = sep
				deleteUser(user)
				fmt.Println(user)
			}

			//fmt.Println(strings.Join(ctx.Args, separator))

			return 0
		},
	}
	cli.AddCommand(addCmd)
	cli.AddCommand(deleteCmd)
	cli.Run()
}
