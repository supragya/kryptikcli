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
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/supragya/kryptik"
)

var generatePrefix string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate RSA 4096 pub/priv key pair for kryptik use",
	Long:  `Generate RSA 4096 pub/priv key pair for kryptik use`,
	Run: func(cmd *cobra.Command, args []string) {
		if generatePrefix == "" {
			log.Error("Give a valid prefix for keyfile names please!")
			return
		}

		log.Info("Generating keys with prefixed name: ", generatePrefix)
		kryptik.GenerateKeysToFiles(generatePrefix)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.PersistentFlags().StringVarP(&generatePrefix, "prefix", "p", "", "GPG RSA 4096 keyfile prefix")
}
