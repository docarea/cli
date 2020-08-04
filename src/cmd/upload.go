/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)



// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload your documentation to docarea",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Args: func(cmd *cobra.Command, args []string) error{

		var message string = ""

		// Hä?
		if len(args) < 1 {
			message += "\nDocumentation path at end of command missing\n"
		}else if len(args) > 1 {
			message += "\nToo many arguments, only documentation path required\n "
		}

		if documentationid == "" || clientid == "" || clientsecret == "" {

			message = "\nPlease specify the following values: \n"

			if documentationid == "" {
				message += "Documentation ID by using --documentation-id [documentation-id]\n"
			}
			if clientid == "" {
				message += "Client ID by using --client-id [client-id]\n"
			}
			if clientsecret == "" {
				message += "Client Secret by using --client-secret [client-secret]\n"
			}

		}

		if message != "" {
			return errors.New(message)
		}

		return nil

	},

	PreRun: func(cmd *cobra.Command, args []string) {

	},

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("upload called\n")

	},
}

// flags for upload command
var documentationid, clientid, clientsecret string

func init() {
	rootCmd.AddCommand(uploadCmd)

	uploadCmd.PersistentFlags().StringVar(&documentationid, "documentation-id",  "", "Documentation ID (required)")
	uploadCmd.PersistentFlags().StringVar(&clientid, "client-id",  "", "Client ID to upload specific documentation (required)")
	uploadCmd.PersistentFlags().StringVar(&clientsecret, "client-secret", "", "Client Secret to upload specific documentation (required)")

}