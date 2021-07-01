package drugtrace

// DataEntTaskDto 
type DataEntTaskDto struct {
    // fileInfoList
    FileInfoList   []string `json:"file_info_list,omitempty" xml:"file_info_list>string,omitempty"`
    // fileNum
    FileNum   int64 `json:"file_num,omitempty" xml:"file_num,omitempty"`
}
