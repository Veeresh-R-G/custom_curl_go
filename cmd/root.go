/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/url"
	"time"

	// "json"

	"os"

	"github.com/Veeresh-R-G/custom_curl_go/httpRequest"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "custom_curl_go",
	Short: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("...")
		verbose, _ := cmd.Flags().GetBool("verbose")
		MethodFlag, _ := cmd.Flags().GetString("request")
		dataFlag, _ := cmd.Flags().GetString("data")
		headers, _ := cmd.Flags().GetStringArray("header")

		if verbose {
			fmt.Println("Verbose description pending ....")
		}
		switch MethodFlag {
		case "GET", "POST", "DELETE":

		default:
			log.Fatalln("Wrong Method Specified")

		}
		u, err := url.Parse(args[1])
		if err != nil {
			log.Fatalln("Couldn't parse given URL")
		}

		HostName := u.Hostname()
		port := u.Port()
		path := u.Path

		if port == "" {
			port = "80"
		}
		fmt.Println("Sending Request ...")
		time.Sleep(2500000000)
		conn, err := httpRequest.HttpTCPConnection(HostName, port)

		if err != nil {
			log.Fatalf("Creating Connection Object : %v\n", err)
		}
		defer conn.Close()

		req, _ := httpRequest.PrepareRequest(MethodFlag, headers, dataFlag, HostName, path)

		_, err = conn.Write([]byte(req))
		if err != nil {
			log.Fatalf("Error while writing to connection Object : %v\n", err)
			return
		}
		buffer := make([]byte, 1024)
		_, err = conn.Read(buffer)
		if err != nil {
			log.Fatalf("Error while writing to buffer : %v\n", err)
			return
		}
		fmt.Printf("Response : \n%v\n", string(buffer))

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//All Flags definition Over here
	var Verbose bool
	var Request string
	var Headers []string
	var Data string

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&Request, "request", "X", "GET", "HTTP request method")
	rootCmd.PersistentFlags().StringVarP(&Data, "data", "d", "", "HTTP request data")
	rootCmd.Flags().StringArrayVarP(&Headers, "header", "H", []string{}, "HTTP request headers")

}
