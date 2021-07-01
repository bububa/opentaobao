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
type TaobaoTanxAuditDepositcreativeAddAPIRequest struct {
    model.Params
    // DSP的memberId
    _memberId   int64
    // dsp用户身份认证的TOKEN
    _token   string
    // 当前时间戳，1970-01-01后的秒数
    _signTime   int64
    // 托管创意信息
    _creative   *CreativeInfoDTO
}

// 初始化TaobaoTanxAuditDepositcreativeAddAPIRequest对象
func NewTaobaoTanxAuditDepositcreativeAddRequest() *TaobaoTanxAuditDepositcreativeAddAPIRequest{
    return &TaobaoTanxAuditDepositcreativeAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTanxAuditDepositcreativeAddAPIRequest) GetApiMethodName() string {
    return "taobao.tanx.audit.depositcreative.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTanxAuditDepositcreativeAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MemberId Setter
// DSP的memberId
func (r *TaobaoTanxAuditDepositcreativeAddAPIRequest) SetMemberId(_memberId int64) error {
    r._memberId = _memberId
    r.Set("member_id", _memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxAuditDepositcreativeAddAPIRequest) GetMemberId() int64 {
    return r._memberId
}
// Token Setter
// dsp用户身份认证的TOKEN
func (r *TaobaoTanxAuditDepositcreativeAddAPIRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoTanxAuditDepositcreativeAddAPIRequest) GetToken() string {
    return r._token
}
// SignTime Setter
// 当前时间戳，1970-01-01后的秒数
func (r *TaobaoTanxAuditDepositcreativeAddAPIRequest) SetSignTime(_signTime int64) error {
    r._signTime = _signTime
    r.Set("sign_time", _signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxAuditDepositcreativeAddAPIRequest) GetSignTime() int64 {
    return r._signTime
}
// Creative Setter
// 托管创意信息
func (r *TaobaoTanxAuditDepositcreativeAddAPIRequest) SetCreative(_creative *CreativeInfoDTO) error {
    r._creative = _creative
    r.Set("creative", _creative)
    return nil
}

// Creative Getter
func (r TaobaoTanxAuditDepositcreativeAddAPIRequest) GetCreative() *CreativeInfoDTO {
    return r._creative
}
