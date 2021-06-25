package scbp

// KeywordInfo 
type KeywordInfo struct {

    // 词
    Word   string `json:"word,omitempty"`

    // 价格
    Price   string `json:"price,omitempty"`

    // 状态
    OnlineStatus   int64 `json:"online_status,omitempty"`

    // 主键id
    Id   int64 `json:"id,omitempty"`

}
