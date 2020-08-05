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
	"docArea/core"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"net/url"
)

// build_docArea.exe upload --documentation-id 720507cb-a770-4e11-8e39-d5ed4d64f681 --client-id iOOhtmvY4w8zwfER7Ls6gfhOKjmsT8x1259Vu4Ob --client-secret ZR0HUciLgJgdQVycVSKtMuJ3AsuPL9b9yHwsUKsdljkXTepnOYc7dDF4uj7fLF4gVtKrQ6skjwTO8T8N7HLKpVr6yy0jR3J5vIpOmrkTZfar4IJJY4JjfgtG8ln0Zvoc path

var api_endpoint = core.Config_api_endpoint
var access_token string

// flags for upload command
var documentationid, clientid, clientsecret string

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
			message += "Documentation path at end of command missing\n"
		}else if len(args) > 1 {
			message += "Too many arguments, only documentation path required\n "
		}

		if documentationid == "" || clientid == "" || clientsecret == "" {

			message += "Please specify the following flags: \n"

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

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("upload called\n")

		response, error := http.PostForm(api_endpoint + "/oauth2/token/", url.Values{
			"grant_type": {"client_credentials"},
			"scope": {"upload_documentation"},
			"client_id": {clientid},
			"client_secret": {clientsecret}})

		if error != nil {
			fmt.Println(error)
			return
		}

		var result map[string]interface{}

		json.NewDecoder(response.Body).Decode(&result)

		access_token := result["access_token"]

		fmt.Println(access_token)

		type uploadrequestbody struct {
			State  string `json:"state"`
			Code   int    `json:"code"`
			Object struct {
				DocumentationID string `json:"documentationId"`
				Size            int    `json:"size"`
				Checksum        string `json:"checksum"`
				SendMeta        bool   `json:"sendMeta"`
			} `json:"object"`
		}



	},
}



func init() {
	rootCmd.AddCommand(uploadCmd)

	uploadCmd.PersistentFlags().StringVar(&documentationid, "documentation-id",  "", "Documentation ID (required)")
	uploadCmd.PersistentFlags().StringVar(&clientid, "client-id",  "", "Client ID to upload specific documentation (required)")
	uploadCmd.PersistentFlags().StringVar(&clientsecret, "client-secret", "", "Client Secret to upload specific documentation (required)")

}
