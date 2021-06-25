package iot

// OpenBaseInfo 
type OpenBaseInfo struct {

    // 账户体系隔离
    Schema   string `json:"schema,omitempty"`

    // 用户ID，此处传入第三方账户体系的用户id
    UserId   string `json:"user_id,omitempty"`

    // 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
    UtdId   string `json:"utd_id,omitempty"`

    // 扩展信息，用于存放APP类型等
    Ext   string `json:"ext,omitempty"`

}
