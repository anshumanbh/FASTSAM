package tasks

import (
	"fmt"
	"strconv"

	"github.com/fsouza/go-dockerclient"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// NmapScan ...
func NmapScan(tstring string) (string, error) {
	envars := []string{tstring}

	client, err := docker.NewClientFromEnv()
	check(err)

	nmapOptions := docker.CreateContainerOptions{

		Config: &docker.Config{
			Image: "dockertools_nmap",
			Env:   envars,
		},
	}

	nmapContainer, err := client.CreateContainer(nmapOptions)
	check(err)

	err = client.StartContainer(nmapContainer.ID, &docker.HostConfig{})
	check(err)

	nmapStatusCode, err := client.WaitContainer(nmapContainer.ID)
	check(err)

	if nmapStatusCode != 0 {
		fmt.Errorf("Process returned bad status-nmap code: %d", nmapStatusCode)
	}

	defer func() {
		if err := client.RemoveContainer(docker.RemoveContainerOptions{
			ID:    nmapContainer.ID,
			Force: true,
		}); err != nil {
			panic(err)
		}
	}()

	ns := strconv.Itoa(nmapStatusCode)
	fs := "Status Code of nmap container is " + ns
	return fs, nil
}

// MasScan ...
func MasScan(tstring string) (string, error) {
	envars := []string{tstring}

	//envars looks like TARGETS=192.168.1.1 192.168.1.2 192.168.1.3

	client, err := docker.NewClientFromEnv()
	check(err)

	masscanOptions := docker.CreateContainerOptions{

		Config: &docker.Config{
			Image: "dockertools_masscan",
			Env:   envars,
		},
	}

	masscanContainer, err := client.CreateContainer(masscanOptions)
	check(err)

	err = client.StartContainer(masscanContainer.ID, &docker.HostConfig{})
	check(err)

	masscanStatusCode, err := client.WaitContainer(masscanContainer.ID)
	check(err)

	if masscanStatusCode != 0 {
		fmt.Errorf("Process returned bad status-masscan code: %d", masscanStatusCode)
	}

	defer func() {
		if err := client.RemoveContainer(docker.RemoveContainerOptions{
			ID:    masscanContainer.ID,
			Force: true,
		}); err != nil {
			panic(err)
		}
	}()

	ms := strconv.Itoa(masscanStatusCode)
	fs := "Status Code of masscan container is " + ms
	return fs, nil
}
