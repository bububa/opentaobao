package servicecenter

// CreditInfoTopDto 
type CreditInfoTopDto struct {

    // 身份证
    IdentityNo   string `json:"identity_no,omitempty"`

    // 手机号
    Mobile   int64 `json:"mobile,omitempty"`

    // 名字
    Name   string `json:"name,omitempty"`

    // 是否通过
    Pass   bool `json:"pass,omitempty"`

    // 被拒原因，只支持传入1,2,3,4.其中1代表综合评分不足；2代表黑名单客群；3代表信用不良；4代表其他
    RejectMsg   string `json:"reject_msg,omitempty"`

    // 唯一id
    Uuid   string `json:"uuid,omitempty"`

    // 额度，单位分，假设离线通过的不用给额度
    Amount   int64 `json:"amount,omitempty"`

    // 0代表已经完成，1代表还需要进一步评估处理
    Flag   int64 `json:"flag,omitempty"`

}
