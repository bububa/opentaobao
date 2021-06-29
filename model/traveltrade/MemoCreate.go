package traveltrade

// MemoCreate 
type MemoCreate struct {
    // 备注添加时间
    Created   string `json:"created,omitempty" xml:"created,omitempty"`
    // 交易ID
    Tid   int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}
