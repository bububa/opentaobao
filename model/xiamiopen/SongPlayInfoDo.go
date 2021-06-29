package xiamiopen

// SongPlayInfoDo 
type SongPlayInfoDo struct {
    // 试听文件列表
    ListenFileList   []ListenFileDo `json:"listen_file_list,omitempty" xml:"listen_file_list>listen_file_do,omitempty"`
    // 歌曲id
    SongId   int64 `json:"song_id,omitempty" xml:"song_id,omitempty"`
}
