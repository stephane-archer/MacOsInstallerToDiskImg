package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func runCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func createImgDisk(installerAppPath string, resultPath string) error {
	const diskImageSize = "7000m"
	const tmpImagePath = "/tmp/MacOSTmpDiskImg.dmg"
	var installerName = path.Base(installerAppPath)
	var mountPath = path.Join("/Volumes/", installerName[:len(installerName)-len(path.Ext(installerName))])
	var installerPath = path.Join(installerAppPath, "Contents/Resources/createinstallmedia")
	err := runCommand("hdiutil", "create", "-o", tmpImagePath, "-size", diskImageSize, "-layout", "SPUD", "-fs", "HFS+J")
	if err != nil {
		return err
	}
	defer func() {
		err := os.Remove(tmpImagePath)
		if err != nil {
			fmt.Println(err)
		}
	}()
	err = runCommand("hdiutil", "attach", tmpImagePath, "-noverify", "-mountpoint", mountPath)
	if err != nil {
		return err
	}
	err = runCommand("sudo", installerPath, "--volume", mountPath, "--nointeraction")
	if err != nil {
		return err
	}
	err = runCommand("hdiutil", "detach", mountPath)
	if err != nil {
		return err
	}
	err = runCommand("hdiutil", "convert", tmpImagePath, "-format", "UDTO", "-o", resultPath)
	if err != nil {
		return err
	}
	err = os.Rename(resultPath+".cdr", resultPath)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var argv = os.Args[1:]
	if len(argv) < 2 {
		fmt.Println("not the right args, first arg is the installer, the second arg is the destination")
		os.Exit(1)
	}
	err := createImgDisk(argv[0], argv[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
