package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
dsp托管创意新增接口 API请求
taobao.tanx.audit.depositcreative.add

dsp托管创意新增接口
*/
type TaobaoTanxAuditDepositcreativeAddRequest struct {
    model.Params
    // DSP的memberId
    memberId   int64
    // dsp用户身份认证的TOKEN
    token   string
    // 当前时间戳，1970-01-01后的秒数
    signTime   int64
    // 托管创意信息
    creative   *CreativeInfoDTO
}

// 初始化TaobaoTanxAuditDepositcreativeAddRequest对象
func NewTaobaoTanxAuditDepositcreativeAddRequest() *TaobaoTanxAuditDepositcreativeAddRequest{
    return &TaobaoTanxAuditDepositcreativeAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTanxAuditDepositcreativeAddRequest) GetApiMethodName() string {
    return "taobao.tanx.audit.depositcreative.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTanxAuditDepositcreativeAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MemberId Setter
// DSP的memberId
func (r *TaobaoTanxAuditDepositcreativeAddRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxAuditDepositcreativeAddRequest) GetMemberId() int64 {
    return r.memberId
}
// Token Setter
// dsp用户身份认证的TOKEN
func (r *TaobaoTanxAuditDepositcreativeAddRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r TaobaoTanxAuditDepositcreativeAddRequest) GetToken() string {
    return r.token
}
// SignTime Setter
// 当前时间戳，1970-01-01后的秒数
func (r *TaobaoTanxAuditDepositcreativeAddRequest) SetSignTime(signTime int64) error {
    r.signTime = signTime
    r.Set("sign_time", signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxAuditDepositcreativeAddRequest) GetSignTime() int64 {
    return r.signTime
}
// Creative Setter
// 托管创意信息
func (r *TaobaoTanxAuditDepositcreativeAddRequest) SetCreative(creative *CreativeInfoDTO) error {
    r.creative = creative
    r.Set("creative", creative)
    return nil
}

// Creative Getter
func (r TaobaoTanxAuditDepositcreativeAddRequest) GetCreative() *CreativeInfoDTO {
    return r.creative
}
