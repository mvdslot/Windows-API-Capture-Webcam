package main

import (
	"fmt"
	"os/exec"
	"syscall"
	"unsafe"
	//"golang.design/x/clipboard"
)

var (
	avicap32   = syscall.NewLazyDLL("avicap32.dll")
	proccapCreateCaptureWindowA  = avicap32.NewProc("capCreateCaptureWindowA")
	
	user32  = syscall.NewLazyDLL("user32.dll")
	procSendMessageA = user32.NewProc("SendMessageA")
)

func CaptureWebcam() {
	// Init returns an error if the package is not ready for use.
	//err := clipboard.Init()
	//if err != nil {
	//	panic(err)
	//}

	var name = "WebcamCapture"
	handle, _, _ := proccapCreateCaptureWindowA.Call(uintptr(unsafe.Pointer(&name)), 0, 0, 0, 320, 240, 0, 0)
	procSendMessageA.Call(handle, 0x40A, 0, 0) //WM_CAP_DRIVER_CONNECT
	procSendMessageA.Call(handle, 0x432, 30, 0) //WM_CAP_SET_PREVIEW
	procSendMessageA.Call(handle, 0x43C, 0, 0) //WM_CAP_GRAB_FRAME
	procSendMessageA.Call(handle, 0x41E, 0, 0) //WM_CAP_EDIT_COPY
	procSendMessageA.Call(handle, 0x40B, 0, 0) //WM_CAP_DRIVER_DISCONNECT

	//img := clipboard.Read(clipboard.FmtImage)

	//_, err = camera.Write(img)
	_, err := exec.Command("PowerShell", "-Command", "Add-Type", "-AssemblyName", fmt.Sprintf("System.Windows.Forms;$clip=[Windows.Forms.Clipboard]::GetImage();if ($clip -ne $null) { $clip.Save('%s') };", "Image.png")).CombinedOutput()

	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	CaptureWebcam()
}
