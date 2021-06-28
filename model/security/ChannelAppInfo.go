package security

// ChannelAppInfo 
/* model for simplify = false
type ChannelAppInfo struct {

    // 渠道名称,多渠道加固才有值
    
    Channel   string `json:"channel,omitempty"`
    

    // 加固后的APP下载地址
    
    AppUrl   string `json:"app_url,omitempty"`
    

}
*/

// ChannelAppInfo 
type ChannelAppInfo struct {

    // 渠道名称,多渠道加固才有值
    Channel   string `json:"channel,omitempty"`

    // 加固后的APP下载地址
    AppUrl   string `json:"app_url,omitempty"`

}
