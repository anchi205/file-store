package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lsCmd)
}

func listCLIHandler() {
	url := "http://localhost:8080/" + "list"

	response, err := http.Get(url)
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

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list files command",
	Long:  "list files command",
	Run: func(cmd *cobra.Command, args []string) {
		listCLIHandler()
	},
}
