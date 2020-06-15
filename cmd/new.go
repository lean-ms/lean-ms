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
	"fmt"
	"html/template"
	"os"
	"path"

	"github.com/lean-ms/lean-ms/utils"
	"github.com/spf13/cobra"
)

type AppConfig struct {
	Name              string
	UnderscoreAppName string
}

func createNewApp(cmd *cobra.Command, args []string) {
	appName := args[0]
	appConfig := AppConfig{appName, utils.Underscore(appName)}
	fmt.Printf("Creating new app named %s", appName)
	createDbConfig(appConfig)
}

func createDbConfig(appConfig AppConfig) {
	dbConfigTempl := template.Must(template.ParseFiles("templates/database.yml"))
	configPath := path.Join(appConfig.Name, "config")
	err := os.MkdirAll(configPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	dbConfigFile, err := os.Create(path.Join(configPath, "database.yml"))
	if err != nil {
		panic(err)
	}
	err = dbConfigTempl.Execute(dbConfigFile, appConfig)
	if err != nil {
		panic(err)
	}
}

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new application",
	Long:  ``,
	Run:   createNewApp,
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}