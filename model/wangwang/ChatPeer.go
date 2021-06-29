package wangwang

// ChatPeer 
type ChatPeer struct {
    // 时间
    Date   string `json:"date,omitempty" xml:"date,omitempty"`
    // 账号。长ID
    Uid   string `json:"uid,omitempty" xml:"uid,omitempty"`
}
