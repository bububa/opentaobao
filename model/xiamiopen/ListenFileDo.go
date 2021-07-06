package xiamiopen

// ListenFileDo 结构体
type ListenFileDo struct {
	// 试听文件地址
	ListenFile string `json:"listen_file,omitempty" xml:"listen_file,omitempty"`
	// 歌曲品质，l为低品质，h为高品质，s为无损
	Quality string `json:"quality,omitempty" xml:"quality,omitempty"`
	// 超时时间
	Expire int64 `json:"expire,omitempty" xml:"expire,omitempty"`
}
