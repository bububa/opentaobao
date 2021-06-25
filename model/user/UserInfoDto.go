package user

// UserInfoDto 
type UserInfoDto struct {

    // 返回数据签名，signature = sha1（raw_data 下所有字段 + appSecret，按字符串升级排列），用于校验关键数据是否被篡改
    Signature   string `json:"signature,omitempty"`

    // 用户相关的原始数据
    RawData   *UserInfoBaseDto `json:"raw_data,omitempty"`

    // 用户openId
    OpenId   string `json:"open_id,omitempty"`

    // 用户结果code
    UserResultCode   *UserResultCode `json:"user_result_code,omitempty"`

}
