package cmns

// CmnsMessage 
type CmnsMessage struct {
    // 应用订阅主题
    Topic   string `json:"topic,omitempty" xml:"topic,omitempty"`
    // 消息是否压缩,0:二进制压缩,1:utf8文本
    Text   int64 `json:"text,omitempty" xml:"text,omitempty"`
    // 消息发送对象
    Receiver   *Receiver `json:"receiver,omitempty" xml:"receiver,omitempty"`
    // 去重码,1-6位整数,0将视为不填,如果填写则同一appKey相同的去重码消息将会被去重，只保留最新的一条，请谨慎使用
    CollapseKey   string `json:"collapse_key,omitempty" xml:"collapse_key,omitempty"`
    // 消息过期时间,lunix时间戳,单位S,为空或不填时为当前时间的1小时后
    Expiration   int64 `json:"expiration,omitempty" xml:"expiration,omitempty"`
    // 消息是否加密,1:加密,0:不加密
    Important   int64 `json:"important,omitempty" xml:"important,omitempty"`
    // 消息内容
    Data   string `json:"data,omitempty" xml:"data,omitempty"`
    // 消息图标
    Ico   string `json:"ico,omitempty" xml:"ico,omitempty"`
    // yunos4.0专用,限制消息只有指定应用来侦听
    Uri   string `json:"uri,omitempty" xml:"uri,omitempty"`
    // 显示方式 1:通知中心 6:后台
    Showtype   int64 `json:"showtype,omitempty" xml:"showtype,omitempty"`
    // 消息标题
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 消息参数,json字符串格式
    Parameter   string `json:"parameter,omitempty" xml:"parameter,omitempty"`
    // 消息优先权
    Priority   int64 `json:"priority,omitempty" xml:"priority,omitempty"`
    // 为action或广播侦听通道，要求填写应用包名(package)
    Program   string `json:"program,omitempty" xml:"program,omitempty"`
    // 响应方式 0:无响应 4:打开应用 7:发送广播
    Responsetype   int64 `json:"responsetype,omitempty" xml:"responsetype,omitempty"`
    // 仅IOS应用推送时使用，默认值为0，开发环境为1，生产环境为0
    DeployStatus   int64 `json:"deploy_status,omitempty" xml:"deploy_status,omitempty"`
}
