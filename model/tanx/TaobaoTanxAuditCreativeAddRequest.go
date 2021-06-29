package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创意预审新增接口 API请求
taobao.tanx.audit.creative.add

创意预审新增接口
*/
type TaobaoTanxAuditCreativeAddRequest struct {
    model.Params
    // DSP的memberId
    _memberId   int64
    // dsp用户身份认证的TOKEN
    _token   string
    // 当前时间戳，1970-01-01后的秒数
    _signTime   int64
    // 预审核创意对象
    _creative   *CreativeParamDTO
}

// 初始化TaobaoTanxAuditCreativeAddRequest对象
func NewTaobaoTanxAuditCreativeAddRequest() *TaobaoTanxAuditCreativeAddRequest{
    return &TaobaoTanxAuditCreativeAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTanxAuditCreativeAddRequest) GetApiMethodName() string {
    return "taobao.tanx.audit.creative.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTanxAuditCreativeAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MemberId Setter
// DSP的memberId
func (r *TaobaoTanxAuditCreativeAddRequest) SetMemberId(_memberId int64) error {
    r._memberId = _memberId
    r.Set("member_id", _memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxAuditCreativeAddRequest) GetMemberId() int64 {
    return r._memberId
}
// Token Setter
// dsp用户身份认证的TOKEN
func (r *TaobaoTanxAuditCreativeAddRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoTanxAuditCreativeAddRequest) GetToken() string {
    return r._token
}
// SignTime Setter
// 当前时间戳，1970-01-01后的秒数
func (r *TaobaoTanxAuditCreativeAddRequest) SetSignTime(_signTime int64) error {
    r._signTime = _signTime
    r.Set("sign_time", _signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxAuditCreativeAddRequest) GetSignTime() int64 {
    return r._signTime
}
// Creative Setter
// 预审核创意对象
func (r *TaobaoTanxAuditCreativeAddRequest) SetCreative(_creative *CreativeParamDTO) error {
    r._creative = _creative
    r.Set("creative", _creative)
    return nil
}

// Creative Getter
func (r TaobaoTanxAuditCreativeAddRequest) GetCreative() *CreativeParamDTO {
    return r._creative
}
