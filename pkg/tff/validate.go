package tff

import "fmt"

func ValidateMain(configPath string) error {
	_, err := LoadYamlFile(configPath)
	if err != nil {
		return err
	}
	fmt.Printf("Config file %q is valid.\n", configPath)
	return nil
}
