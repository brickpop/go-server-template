package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	log "go.vocdoni.io/dvote/log"
)

// Init performs the initial viper setup
func Init(rootCmd *cobra.Command) {
	// Bind config to CLI params
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("cert", rootCmd.PersistentFlags().Lookup("cert"))
	viper.BindPFlag("key", rootCmd.PersistentFlags().Lookup("key"))
	viper.BindPFlag("tls", rootCmd.PersistentFlags().Lookup("tls"))

	viper.AddConfigPath(".")

	configFile := viper.GetString("config")
	viper.SetConfigFile(configFile)
	// viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	}
}

// DefineCliFlags defines the flags accepted by the CLI
func DefineCliFlags(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().String("config", "", "the config file to use")
	rootCmd.PersistentFlags().String("cert", "", "the certificate file (TLS only)")
	rootCmd.PersistentFlags().String("key", "", "the TLS encryption key file")
	rootCmd.PersistentFlags().Bool("tls", false, "whether to use TLS encryption (cert and key required)")
	rootCmd.PersistentFlags().IntP("port", "p", 8080, "port to bind to")
}
