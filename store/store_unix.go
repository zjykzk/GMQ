package store

import (
	"syscall"
)

func mmap(fd, sz int) ([]byte, error) {
	return syscall.Mmap(fd, 0, sz, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
}

func unmap(addr, len uintptr) error {
	_, _, errno := syscall.Syscall(syscall.SYS_MUNMAP, addr, len, 0)
	if errno != 0 {
		return syscall.Errno(errno)
	}
	return nil
}

func flush(addr, len uintptr) error {
	_, _, errno := syscall.Syscall(syscall.SYS_MSYNC, addr, len, syscall.MS_SYNC)
	if errno != 0 {
		return syscall.Errno(errno)
	}
	return nil
}
