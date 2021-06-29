package nrt

// TmallNrtItemMainSynchronizeResultDO 
type TmallNrtItemMainSynchronizeResultDO struct {
    // 返回值
    Data   *NrtItemSyncResultDTO `json:"data,omitempty" xml:"data,omitempty"`
    // 调用是否成功
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`
}
