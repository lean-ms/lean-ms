/*
Copyright © 2020 PAULO SOARES <phsoares.ita@gmail.com>

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
	"log"
	"os"
	"path"

	"github.com/lean-ms/lean-ms/cmd/helpers"
	"github.com/lean-ms/utils"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate resources",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		resource := args[0]
		switch resource {
		case "scaffold":
			resourceName := args[1]
			createResource(resourceName, args[2:]...)
		default:
			log.Fatalf("Use scaffold as resource. Used: %s", resource)
		}
	},
}

type ModelTemplate struct {
	camelizedModelName string
	properties         []string
}

func createResource(resourceName string, params ...string) {
	templatePath := path.Join(helpers.GetBasePath(), "templates")
	modelsTemplateFile := path.Join(templatePath, "models", "model.go")
	modelsTemplate := template.Must(template.ParseFiles(modelsTemplateFile))
	err := os.MkdirAll("models", os.ModePerm)
	if err != nil {
		panic(err)
	}

	modelTemplate := &ModelTemplate{utils.Camelize(resourceName), make([]string, 0)}

	modelFilepath := fmt.Sprintf("%s.go", utils.Underscore(resourceName))
	modelFile, err := os.Create(path.Join("models", modelFilepath))
	if err != nil {
		panic(err)
	}
	err = modelsTemplate.Execute(modelFile, modelTemplate)
	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
