package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创意修改接口 API请求
taobao.tanx.audit.creative.modify

创意修改接口
*/
type TaobaoTanxAuditCreativeModifyRequest struct {
    model.Params
    // DSP用户ID
    _memberId   int64
    // dsp用户身份认证的TOKEN
    _token   string
    // 当前时间戳，1970-01-01后的秒数
    _signTime   int64
}

// 初始化TaobaoTanxAuditCreativeModifyRequest对象
func NewTaobaoTanxAuditCreativeModifyRequest() *TaobaoTanxAuditCreativeModifyRequest{
    return &TaobaoTanxAuditCreativeModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTanxAuditCreativeModifyRequest) GetApiMethodName() string {
    return "taobao.tanx.audit.creative.modify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTanxAuditCreativeModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MemberId Setter
// DSP用户ID
func (r *TaobaoTanxAuditCreativeModifyRequest) SetMemberId(_memberId int64) error {
    r._memberId = _memberId
    r.Set("member_id", _memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxAuditCreativeModifyRequest) GetMemberId() int64 {
    return r._memberId
}
// Token Setter
// dsp用户身份认证的TOKEN
func (r *TaobaoTanxAuditCreativeModifyRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoTanxAuditCreativeModifyRequest) GetToken() string {
    return r._token
}
// SignTime Setter
// 当前时间戳，1970-01-01后的秒数
func (r *TaobaoTanxAuditCreativeModifyRequest) SetSignTime(_signTime int64) error {
    r._signTime = _signTime
    r.Set("sign_time", _signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxAuditCreativeModifyRequest) GetSignTime() int64 {
    return r._signTime
}
