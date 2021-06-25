package security

// JaqResourceResult 
type JaqResourceResult struct {

    // 请求事件唯一标识
    EventId   string `json:"event_id,omitempty"`

    // 资源包的md5
    Md5   string `json:"md5,omitempty"`

    // 资源的cdn下载链接
    Url   string `json:"url,omitempty"`

    // 资源版本号
    Version   string `json:"version,omitempty"`

}
