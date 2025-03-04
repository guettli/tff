package tff

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/holoplot/go-evdev"
)

type CombosCmdConfig struct {
	Debug       bool
	DevicePaths []string
	ConfigFile  string
	Logfile     string
}

type device struct {
	id int

	path string

	// sourceDev is the device we read from
	sourceDev *evdev.InputDevice

	// outDev is the device we write to. Created via CloneDevice
	outDev *evdev.InputDevice
}

func (d *device) Open() error {
	sourceDev, err := evdev.Open(d.path)
	if err != nil {
		return fmt.Errorf("failed to open device %q: %w", d.path, err)
	}

	err = sourceDev.Grab()
	if err != nil {
		return fmt.Errorf("failed to grab device: %w", err)
	}
	outDev, err := evdev.CloneDevice(fmt.Sprintf("tff-clone-%d-%d", os.Getpid(), d.id), sourceDev)
	if err != nil {
		return errors.Join(sourceDev.Close(), err)
	}
	d.sourceDev = sourceDev
	d.outDev = outDev
	return nil
}

func CombosMain(ctx context.Context, cmdconfig CombosCmdConfig) error {
	combos, err := LoadYamlFile(cmdconfig.ConfigFile)
	if err != nil {
		return err
	}

	if len(cmdconfig.DevicePaths) == 0 {
		p, err := findDev()
		if err != nil {
			return err
		}

		idPath, err := getDeviceAlias(p)
		alias := "(no alias found for device)"
		if err == nil {
			alias = fmt.Sprintf("(%s)", idPath)
		}
		fmt.Printf("%s %s %q\n", usingDeviceMessage, alias, p)
		cmdconfig.DevicePaths = []string{p}
	}

	devices := make([]*device, 0, len(cmdconfig.DevicePaths))
	good := 0
	openErrors := make([]error, 0, len(cmdconfig.DevicePaths))
	for i, path := range cmdconfig.DevicePaths {
		dev := device{
			id:   i,
			path: path,
		}
		devices = append(devices, &dev)
		err := dev.Open()
		if err != nil {
			openErrors = append(openErrors, err)
			continue
		}
		good++
	}

	if good == 0 {
		return fmt.Errorf("no devices could be opened: %w", errors.Join(openErrors...))
	}
	ctx, cancel := context.WithCancelCause(ctx)
	errorChannel := make(chan error)
	for i := range devices {
		go handleOneDevice(ctx, combos, devices[i], errorChannel)
	}
	err = <-errorChannel
	fmt.Printf("error, stopping now: %v\n", err)
	cancel(err)
	for range len(cmdconfig.DevicePaths) - 1 {
		err := <-errorChannel
		fmt.Println(err)
	}
	return nil
}
