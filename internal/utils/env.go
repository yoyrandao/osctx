package utils

import "os"

func GetOSCloud() string {
	cloud := os.Getenv("OS_CLOUD")
	if cloud == "" {
		cloud = "<none>"
	}

	return cloud
}
