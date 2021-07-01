package nrt

// TmallNrtItemMainSynchronizeResultDo 
type TmallNrtItemMainSynchronizeResultDo struct {
    // 返回值
    Data   *NrtItemSyncResultDto `json:"data,omitempty" xml:"data,omitempty"`
    // 调用是否成功
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`
}
