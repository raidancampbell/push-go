package cmd

import (
	"fmt"
	"github.com/gregdel/pushover"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "push-go",
	Short: "A CLI for sending Pushover push notifications to your device",
	Long: "A CLI for sending Pushover push notifications to your device.  A config file should exist at $HOME/.config/.push-go.yaml containing PUSHOVER_KEY and PUSHOVER_RECIPIENT. env overrides respected",
	Run: func(cmd *cobra.Command, args []string) {
		key := viper.GetString("PUSHOVER_KEY")
		if key == "" {
			logrus.Fatal("configuration required: PUSHOVER_KEY not found in configuration file")
		}
		recipient := viper.GetString("PUSHOVER_RECIPIENT")
		if recipient == "" {
			logrus.Fatal("configuration required: PUSHOVER_RECIPIENT not found in configuration file")
		}

		app := pushover.New(key)
		recip := pushover.NewRecipient(recipient)

		message, err := app.SendMessage(&pushover.Message{
			Message: strings.Join(args, ""),
			Title:   "push-go CLI",
		}, recip)
		if err != nil {
			panic(err)
		}
		logrus.Infof("%+v", message)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/.push-go.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(fmt.Sprintf("%s%c%s",home, os.PathSeparator, ".config"))
		viper.SetConfigName(".push-go")
	}
	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logrus.Infof("Using config file: %s", viper.ConfigFileUsed())
	}
}
