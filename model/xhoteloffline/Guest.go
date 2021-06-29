package xhoteloffline

// Guest 
type Guest struct {

    // 是否主入住人，该入住人会参与信用住结算扣款，多个入住人时必须有且仅有一个该字段设置为true
    
    IsMain   bool `json:"is_main,omitempty" xml:"is_main,omitempty"`
    

    // 加密方式, 默认0: 不加密, 信息会通过淘宝开放平台传输, 阿里旅行可以获取到具体信息;  1: SHA-1不可逆加密,  阿里旅行方面无法解析到具体信息, 只用于做信息匹配.注意加密后生成40字节长度的字符串
    
    EncryptType   int64 `json:"encrypt_type,omitempty" xml:"encrypt_type,omitempty"`
    

    // 姓名, 如果加密方式设置为1, 传入加密后的姓名
    
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    

    // 证件类型, 默认0:身份证; 1: 护照; 2:警官证; 3:士兵证; 4: 回乡证
    
    IdType   int64 `json:"id_type,omitempty" xml:"id_type,omitempty"`
    

    // 证件号, 如果加密方式设置为1, 传入加密后的证件号
    
    IdNumber   string `json:"id_number,omitempty" xml:"id_number,omitempty"`
    

    // 手机号, 如果加密方式设置为1, 传入加密后的手机号
    
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    

}
