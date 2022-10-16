package main

import "fmt"

func arrays() {
	var DeploymentOptions = [4]string{"R-pi", "AWS", "GCP", "Azure"}

	for i := 0; i < len(DeploymentOptions); i++ {
		option := DeploymentOptions[i]
		fmt.Println(i, option)
	}
}

func differentloop() {
	DeploymentOptions := [...]string{"R-pi", "AWS", "GCP", "Azure"}

	for index, option := range DeploymentOptions {
		fmt.Println(index, option)
	}
}
