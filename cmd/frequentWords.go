package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(freqWordsCmd)
	freqWordsCmd.Flags().StringP("order", "s", "dsc", "sort order")
	freqWordsCmd.Flags().StringP("limit", "l", "10", "limit")
}

func freqWordsCLIHandler(baseURL string) {
	response, err := http.Get(baseURL)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(string(body))
}

var freqWordsCmd = &cobra.Command{
	Use:   "freq-words",
	Short: "frequent words cmd",
	Long:  "frequent words cmd",
	Run: func(cmd *cobra.Command, args []string) {
		sortorder, _ := cmd.Flags().GetString("order")
		limit, _ := cmd.Flags().GetString("limit")
		baseURL := os.Getenv("BASE_URL") + "/frequentWords"
		if sortorder != "" {
			baseURL = baseURL + "?sortorder=" + sortorder
		}
		if limit != "" {
			baseURL = baseURL + "&limit=" + limit
		}
		freqWordsCLIHandler(baseURL)
	},
}
