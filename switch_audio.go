package audio

import (
	"os/exec"
	"strings"
)

func GetCurrentAudioSource() (string, error) {
	cmd := exec.Command("SwitchAudioSource", "-c")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	deviceName := strings.Split(string(output), "\n")[0]
	return deviceName, nil
}

func GetAllAudioSources() ([]string, error) {
	cmd := exec.Command("SwitchAudioSource", "-a")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var audioSources []string
	elements := strings.Split(string(output), "\n")
	for _, deviceName := range elements {
		if deviceName != "" {
			audioSources = append(audioSources, deviceName)
		}
	}
	return audioSources, nil
}

func SetOutputSource(sourceName string) error {
	cmd := exec.Command("SwitchAudioSource", "-s", sourceName)
	_, err := cmd.Output()
	// TODO: handle possible error

	return err
}
