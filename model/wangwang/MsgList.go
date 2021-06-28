package wangwang

// MsgList 
/* model for simplify = false
type MsgList struct {

    // 当direction=1有效，（关键词，数量）列表
    
    WordLists  struct {
        WordCountList  []WordCountList `json:"word_count_list,omitempty"`
    } `json:"word_lists,omitempty"`
    

    // 当direction=1有效，url列表
    
    UrlLists  struct {
        UrlList  []UrlList `json:"url_list,omitempty"`
    } `json:"url_lists,omitempty"`
    

    // 表示消息方向<br/>0:from_id->to_id<br/>1:to_id->from_id
    
    Direction   int64 `json:"direction,omitempty"`
    

    // utc时间
    
    Time   int64 `json:"time,omitempty"`
    

    // 消息类型
    
    Type   int64 `json:"type,omitempty"`
    

    // 当direction=1有效，消息长度
    
    Length   int64 `json:"length,omitempty"`
    

    // 当direction=0有效，完整消息内容
    
    Content   string `json:"content,omitempty"`
    

}
*/

// MsgList 
type MsgList struct {

    // 当direction=1有效，（关键词，数量）列表
    WordLists   []WordCountList `json:"word_lists,omitempty"`

    // 当direction=1有效，url列表
    UrlLists   []UrlList `json:"url_lists,omitempty"`

    // 表示消息方向<br/>0:from_id->to_id<br/>1:to_id->from_id
    Direction   int64 `json:"direction,omitempty"`

    // utc时间
    Time   int64 `json:"time,omitempty"`

    // 消息类型
    Type   int64 `json:"type,omitempty"`

    // 当direction=1有效，消息长度
    Length   int64 `json:"length,omitempty"`

    // 当direction=0有效，完整消息内容
    Content   string `json:"content,omitempty"`

}
