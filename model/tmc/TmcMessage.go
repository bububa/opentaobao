package tmc

// TmcMessage 
/* model for simplify = false
type TmcMessage struct {

    // 消息ID
    
    Id   int64 `json:"id,omitempty"`
    

    // 消息所属主题
    
    Topic   string `json:"topic,omitempty"`
    

    // 消息发布者的AppKey
    
    PubAppKey   string `json:"pub_app_key,omitempty"`
    

    // 消息发布时间
    
    PubTime   string `json:"pub_time,omitempty"`
    

    // 用户的昵称
    
    UserNick   string `json:"user_nick,omitempty"`
    

    // 消息详细内容，格式为JSON/XML
    
    Content   string `json:"content,omitempty"`
    

    // 消息所属的用户编号
    
    UserId   int64 `json:"user_id,omitempty"`
    

}
*/

// TmcMessage 
type TmcMessage struct {

    // 消息ID
    Id   int64 `json:"id,omitempty"`

    // 消息所属主题
    Topic   string `json:"topic,omitempty"`

    // 消息发布者的AppKey
    PubAppKey   string `json:"pub_app_key,omitempty"`

    // 消息发布时间
    PubTime   string `json:"pub_time,omitempty"`

    // 用户的昵称
    UserNick   string `json:"user_nick,omitempty"`

    // 消息详细内容，格式为JSON/XML
    Content   string `json:"content,omitempty"`

    // 消息所属的用户编号
    UserId   int64 `json:"user_id,omitempty"`

}
