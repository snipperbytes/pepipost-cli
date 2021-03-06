package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
	"pepipost-cli/internal/smtpmail"
	"pepipost-cli/internal/apimail"
	"pepipost-cli/internal/reports"
)

var (
	cmds []cli.Command
	author []cli.Author
)

func init(){

	author = []cli.Author{
		{"Vikram Sahu", "vikram@pepipost.com"},
		{"DX Team", "dx@pepipost.com"},
	}
	cmds = []cli.Command{
		{
			Name:        "smtpEmail",
			Usage:       "Send email using SMTP credentials of any Email Service Provider",
			Description: "Used to send Email through SMTP protocol. We recommend to set your smtpusername and smtppassword as enviroment variable for security reason",
			Flags: smtpmail.Getflags(),
			Action: func(c *cli.Context) {
				status,errcode,err_response := smtpmail.Sendmail(c)
				if errcode == 2 {
					fmt.Print(status)
					return
				}
				if err_response != nil { status = "Error" } else { status = "Success" }
				fmt.Println("Status of Email 	:: ",status)
				fmt.Println("Error Response of Email :: ",err_response)
				return
			},
		},
		{
			Name:        "apiEmail",
			Usage:       "Send email using APIKey of Pepipost",
			Description: "Used for Sending email using your Pepipost API key. we recommend to set your apikey as envviroment. by using set APIKEY=063a51211ec334ece4dd0db3f668c63a",
			Flags: apimail.Getflags(),
			Action: func(c *cli.Context) {
				status,errstatus := apimail.Sendmail(c)
				errorstatus := "Success"
				if errstatus == 3 { 
					errorstatus = "Error" 
					fmt.Print(status)
					return
				}
				fmt.Println("Sending Status :: ",errorstatus)
				fmt.Println("Status of API call    :: ",status)
				return
			},
		},
		{
			Name:        "fetchLogs",
			Usage:       "Fetches realtime logs for emails sent using Pepipost",
			Description: "Helps to retrieve real-time report for emails sent using Pepipost",
			Flags: reports.Getflags(),
			Action: func(c *cli.Context) {
				reports.Fetchlogs(c)
			},
		},

	}
}


func main(){
	if len(os.Args) == 0  {
		fmt.Print(`pepipost : missing OPTIONS
Usage : Pepipost [COMMANDS]... [OPTIONS]... 

Try pepipost --help`)
	}else {
		runPepi()
		return
	}
}

func runPepi(){

	app := cli.NewApp()
	app.Name = "pepipost"
	app.Usage = "AI-Powered Email Delivery Engine helps inboxing your business critical emails "
	app.Version = "1.0.0, For linux"
	app.Description = "Pepipost now allows sending email,fetching stats,real-time reports using command line interface. Use below instructions to send email using ‘Pepipost‘ command."
	app.Authors = author
	app.Copyright = "(c) 2019 Pepipost"
	app.Commands = cmds
	app.Action = noArgs
	app.Run(os.Args)

}


func noArgs(c *cli.Context) error {
	return cli.NewExitError(`pepipost: missing COMMAND 
Usage: pepipost [COMMAND]... [ARGUMENT]...

Try pepipost --help for more options.`, 2)
}


