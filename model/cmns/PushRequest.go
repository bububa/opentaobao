package cmns

// PushRequest 
type PushRequest struct {

    // 设备类型，1：android or yunos，2：ios,当deviceType不指定时，服务器端对msg json字段按照android格式解析
    DeviceType   int64 `json:"device_type,omitempty"`

    // 消息过期时间,相对时间(即多少秒后过期,最多7天),单位S,不填默认1小时过期
    Expiration   int64 `json:"expiration,omitempty"`

    // 仅IOS应用推送时使用，0代表生成环境，1代表沙箱环境，默认值为0
    IosEnvironment   int64 `json:"ios_environment,omitempty"`

    // 消息内容, 为json字符串,格式详见http://open.yunos.com/doc/detail?spm=a2c01.7698725.0002.40.ZNPFOJ&documentId=102975
    Msg   string `json:"msg,omitempty"`

    // 消息发送对象
    Receiver   *Receiver `json:"receiver,omitempty"`

    // 消息类型，当前type只能取1，即只能发送给应用透传消息
    Type   int64 `json:"type,omitempty"`

    // 消息发送优先级，范围为1-5，数字越高，优先级越大，不设置默认优先级为2
    Priority   int64 `json:"priority,omitempty"`

    // 业务应用appKey,top调用不需要传此参数
    BizAppKey   string `json:"biz_app_key,omitempty"`

    // 去重码，为应用自己维护的消息唯一标记
    CollapseKey   string `json:"collapse_key,omitempty"`

}
