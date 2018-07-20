// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/hr1sh1kesh/vegeta-wp-load/src"
	"github.com/spf13/cobra"
)

var rate int
var duration int
var user string
var password string

var loadgenCmd = &cobra.Command{
	Use:   "loadgen",
	Short: "Generate Load for Wordpress",
	Long:  `Generate load and exercise Wordpress using Vegeta`,
	Run:   loadGenerator,
}

func init() {
	rootCmd.AddCommand(loadgenCmd)
	loadgenCmd.PersistentFlags().StringVarP(&api, "api", "a", "", "provide the API endpoint to be load tested")
	loadgenCmd.MarkFlagRequired("api")
	loadgenCmd.PersistentFlags().IntVarP(&rate, "rate", "n", 1, "Request Rate per second")
	loadgenCmd.MarkFlagRequired("rate")
	loadgenCmd.PersistentFlags().IntVarP(&duration, "duration", "d", 1, "Duration in seconds")
	loadgenCmd.MarkFlagRequired("duration")
	loadgenCmd.PersistentFlags().StringVarP(&user, "user", "u", "", "Username")
	loadgenCmd.MarkFlagRequired("user")
	loadgenCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Password")
	loadgenCmd.MarkFlagRequired("password")
}

func loadGenerator(cmd *cobra.Command, args []string) {

	creds := user + ":" + password
	src.GenerateLoadData(rate, duration, api, creds)
}
