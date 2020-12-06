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
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/supragya/kryptik"
)

var keyPrefix, message, remote string

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send kryptik messages to a remote endpoint",
	Long:  `Send kryptik messages to a remote endpoint`,
	Run: func(cmd *cobra.Command, args []string) {
		if message == "" || keyPrefix == "" {
			log.Error("Invalid arguments")
			return
		}
		log.Info("Encapsulating message [", message, "] with a signature")
		signedMessage, err := kryptik.GetSignedMessage(keyPrefix+".privkey", keyPrefix+".pubkey", []byte(message))
		if err != nil {
			log.Error("Error while signing message: ", err)
			return
		}
		log.Info("Kryptic message (signed message) generated successfully!")

		log.Info("Sending signed message to remote")

		j, err := json.Marshal(signedMessage)
		if err != nil {
			log.Error("Signed message was not json encodable")
			return
		}

		req, err := http.NewRequest("POST", remote, bytes.NewBuffer(j))
		if err != nil {
			log.Error("Error while making post request: ", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			log.Error("Error while reading post response: ", err)
			return
		}
		defer resp.Body.Close()
		log.Info("Remote response Status:", resp.Status)
		log.Info("Remote response Headers:", resp.Header)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error("Error while reading post response: ", err)
			return
		}
		log.Info("Remote response Body:", string(body))
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.PersistentFlags().StringVarP(&message, "message", "m", "", "Message to send to remote")
	sendCmd.PersistentFlags().StringVarP(&keyPrefix, "prefix", "p", "", "GPG RSA 4096 keyfile prefix (both pubkey, privkey needed)")
	sendCmd.PersistentFlags().StringVarP(&remote, "remote", "r", "", "Remote host to send signed message to")
}
