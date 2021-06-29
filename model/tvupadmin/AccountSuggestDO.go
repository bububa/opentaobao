package tvupadmin

// AccountSuggestDO 
type AccountSuggestDO struct {
    // 用户ID
    Uid   int64 `json:"uid,omitempty" xml:"uid,omitempty"`
    // 优酷账号ID
    Ytid   int64 `json:"ytid,omitempty" xml:"ytid,omitempty"`
    // 用户昵称
    Nick   string `json:"nick,omitempty" xml:"nick,omitempty"`
    // 优酷账号昵称
    YtNick   string `json:"yt_nick,omitempty" xml:"yt_nick,omitempty"`
}
