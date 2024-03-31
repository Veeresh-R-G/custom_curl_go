/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	// "json"
	"log"
	"net"
	"net/url"
	"os"
	"strings"
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
		protocols["DELETE"] = "yes"

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
		if method == "GET" || method == "DELETE" {

			fmt.Fprintf(conn, "%s %s HTTP/1.0\r\nHost: %s\r\n\r\n", method, path, host)
			n, err := conn.Read(buff)
			fmt.Println("Here")
			if err != nil {
				log.Fatalf("Buffer error : %s", err)
			}

			fmt.Printf("\nResponse : \n\n%s", string(buff[:n]))
		} else {
			//For POST Request
			log.Printf("HostName = %v Port = %v Path = %v\n", host, port, path)
			fmt.Println("POST request")
			// Parse postData JSON string
			var data map[string]interface{}
			json.Unmarshal([]byte(postData), &data)
			fmt.Println("POST Data:", data)

			// Parse headers
			headerMap := make(map[string]string)
			for _, header := range headers {
				parts := strings.SplitN(header, ":", 2)
				if len(parts) != 2 {
					fmt.Println("Invalid header format:", header)
					continue
				}
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				headerMap[key] = value
			}
			fmt.Println("Headers:", headerMap)

			// Compose HTTP request
			request := fmt.Sprintf("%s %s HTTP/1.0\r\nHost: %s\r\n", method, path, host)
			for key, value := range headerMap {
				request += fmt.Sprintf("%s: %s\r\n", key, value)
			}
			request += fmt.Sprintf("Content-Type: application/json\r\nContent-Length: %d\r\n\r\n%s", len(postData), postData)

			// Send HTTP request
			fmt.Fprintf(conn, request)

			// Read response
			n, err := conn.Read(buff)
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
	//All Flags definition Over here
	rootCmd.Flags().BoolP("curl_flag", "X", false, "Compulsory Flag")

}
