/*
Copyright © 2025 PH alvesph1@gmail.com
*/
package cmd

import (
	"github.com/alvesph/challenge-stress-test/internal/service"
	"github.com/spf13/cobra"
)

// stressCmd represents the stress command
var stressCmd = &cobra.Command{
	Use:   "stress",
	Short: "Executa teste de carga em uma URL via CLI",
	Long: `Realiza testes de carga em serviços web via linha de comando.

Permite configurar a URL alvo, o número total de requisições e o nível de concorrência.
Ao final, gera um relatório com tempo total, quantidade de requisições bem-sucedidas, erros e distribuição dos códigos de status HTTP.

Exemplo de uso:
  stress --url=https://httpbin.org/get --requests=100 --concurrency=10`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")

		if url == "" {
			cmd.Println("Informe a URL para o teste de estresse")
			return
		}
		if requests <= 0 || concurrency <= 0 {
			cmd.Println("Número de requisições e concorrência devem ser maiores que zero")
			return
		}

		cmd.Printf("Iniciando teste de estresse em %s com %d requisições e %d concorrência\n", url, requests, concurrency)
		service.RunStressTest(url, requests, concurrency, cmd.Println)
	},
}

func init() {
	rootCmd.AddCommand(stressCmd)
	stressCmd.Flags().StringP("url", "u", "", "URL to stress test")
	stressCmd.Flags().IntP("requests", "r", 100, "Number of requests to send")
	stressCmd.Flags().IntP("concurrency", "c", 10, "Number of concurrent requests")
}
