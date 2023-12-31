package stateinf

import (
	"net/http"
	//"strings"
	"syscall"
	"unsafe"
	//"fmt"
)

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func get(w http.ResponseWriter, r *http.Request) {
	//str := strings.Split(r.URL.EscapedPath(), "/")[2]
	//fmt.Printf("%s",str)
	disk := DiskUsage("./")
	Len := unsafe.Sizeof(*disk)
	bytedisk := &SliceMock{
		addr: uintptr(unsafe.Pointer(disk)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(bytedisk))
	//fmt.Printf("%d",data)
	w.Write(data)
	return
}
func DiskUsage(path string) (disk *DiskStatus) {
	var resptr DiskStatus
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	resptr.All = fs.Blocks * uint64(fs.Bsize)
	resptr.Free = fs.Bfree * uint64(fs.Bsize)
	resptr.Used = resptr.All - resptr.Free
	return &resptr
}
