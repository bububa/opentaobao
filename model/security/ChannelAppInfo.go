package security

// ChannelAppInfo 
type ChannelAppInfo struct {
    // 渠道名称,多渠道加固才有值
    Channel   string `json:"channel,omitempty" xml:"channel,omitempty"`
    // 加固后的APP下载地址
    AppUrl   string `json:"app_url,omitempty" xml:"app_url,omitempty"`
}
