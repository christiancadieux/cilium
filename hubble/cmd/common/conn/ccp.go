package conn

import (
	"bufio"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func CcpSetHubbleService(ctx context.Context, vp *viper.Viper) error {

	cluster, err := getCurrentContext()
	if err != nil {
		return err
	}
	hubbleNode, port, err := FindHubbleService(cluster, ctx)
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	vp.Set("server", fmt.Sprintf("%s:%d", hubbleNode, port))
	return nil
}

func getCurrentContext() (string, error) {
	homedir, err := os.UserHomeDir()
	filename := homedir + "/.kube/config"
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("reading kubeconfig %s - %v", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "current-context:") {
			parts := strings.Split(line, " ")
			return parts[1], nil
		}
	}
	return "", fmt.Errorf("CurrentContext not found in %s", filename)
}
