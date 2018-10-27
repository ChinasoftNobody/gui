package gui

import (
	"syscall"
)

func Window()  {
	handle,err :=syscall.LoadLibrary("	User32.dll")
	if err!= nil{
		return
	}
	defer syscall.FreeLibrary(handle)
	proc,err :=syscall.GetProcAddress(handle, "GetVersion")
	if err!=nil{
		return
	}
	r,_,_:=syscall.Syscall(uintptr(proc),0,0,0,0)
	v:=uint32(r)
	major := byte(v)
	minor := uint8(v >> 8)
	build := uint16(v >> 16)
	print("windows version ", major, ".", minor, " (Build ", build, ")\n")
}
