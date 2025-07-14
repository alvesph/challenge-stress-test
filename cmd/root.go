/*
Copyright © 2025 PH alvesph1@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "challenge-stress-test",
	Short: "Executa teste de carga em uma URL via CLI",
	Long: `Realiza testes de carga em serviços web via linha de comando.

Permite configurar a URL alvo, o número total de requisições e o nível de concorrência.
Ao final, gera um relatório com tempo total, quantidade de requisições bem-sucedidas, erros e distribuição dos códigos de status HTTP.

Exemplo de uso:
  stress --url=https://httpbin.org/get --requests=100 --concurrency=10`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.challenge-stress-test.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
