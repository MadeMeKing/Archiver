package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Run:   pack,
}

const packetExtension = "vlc"

var ErrEmptyPath = errors.New("path to file is not specified")

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		handleErr(ErrEmptyPath)
	}

	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}

	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	packed := ""
	fmt.Println(string(data))

	err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}
}

func packedFileName(path string) string {
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packetExtension
}

func init() {
	packCmd.AddCommand(vlcCmd)
}
