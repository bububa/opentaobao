package wdk

// ChannelProp 
type ChannelProp struct {
    // 渠道类型：txd淘鲜达，elm饿了么，shareStore共享库存
    ChannelType   string `json:"channel_type,omitempty" xml:"channel_type,omitempty"`
    // 渠道属性，取值为key-value键值对形式
    Props   []PropField `json:"props,omitempty" xml:"props>prop_field,omitempty"`
}
