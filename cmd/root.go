/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "custom_curl_go",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Sending Request ...")
		time.Sleep(2500000000)

		protocols := make(map[string]interface{})
		protocols["GET"] = 1
		protocols["POST"] = "yes"

		if args[0] != "carlo" {
			log.Fatalln("Wrong Command Used")
		}

		if _, ok := protocols[args[1]]; !ok {
			log.Fatalln("Wrong Protocol Mentioned")
		}
		method := args[1]
		u, err := url.Parse(args[2])

		if err != nil {
			log.Fatalf("Error : %v\n", err)
		}

		host := u.Hostname()
		port := u.Port()
		path := u.Path
		// log.Printf("HostName = %v Port = %v Path = %v\n", host, port, path)

		if port == "" {
			port = "80"
		}

		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))

		fmt.Println("Request Sent ...")

		if err != nil {
			log.Fatalf("TCP Connection Error : %v", err)
		}
		defer conn.Close()

		buff := make([]byte, 1024)
		if method == "GET" {

			fmt.Fprintf(conn, "%s %s HTTP/1.0\r\nHost: %s\r\n\r\n", method, path, host)
			n, err := conn.Read(buff)
			fmt.Println("Here")
			if err != nil {
				log.Fatalf("Buffer error : %s", err)
			}

			fmt.Printf("\nResponse : \n\n%s", string(buff[:n]))
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.custom_curl_go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("curl", "X", false, "This will trigger custom curl")
	// rootCmd.Flags().Bool("X", true, "This will trigger custom curl")
}
