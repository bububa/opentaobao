package yunos

// OperBatchMsg 
type OperBatchMsg struct {

    // 服务号唯一编号
    AppId   string `json:"app_id,omitempty"`

    // YUNOS开放平台模板ID
    TplId   string `json:"tpl_id,omitempty"`

    // 消息体的JSON串，以小说为例
    Content   string `json:"content,omitempty"`

    // 是否推荐消息标识，默认：false
    TjFlag   bool `json:"tj_flag,omitempty"`

}
