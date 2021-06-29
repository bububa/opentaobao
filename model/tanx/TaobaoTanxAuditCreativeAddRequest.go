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
    memberId   int64
    // dsp用户身份认证的TOKEN
    token   string
    // 当前时间戳，1970-01-01后的秒数
    signTime   int64
    // 预审核创意对象
    creative   *CreativeParamDTO
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
func (r *TaobaoTanxAuditCreativeAddRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxAuditCreativeAddRequest) GetMemberId() int64 {
    return r.memberId
}
// Token Setter
// dsp用户身份认证的TOKEN
func (r *TaobaoTanxAuditCreativeAddRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r TaobaoTanxAuditCreativeAddRequest) GetToken() string {
    return r.token
}
// SignTime Setter
// 当前时间戳，1970-01-01后的秒数
func (r *TaobaoTanxAuditCreativeAddRequest) SetSignTime(signTime int64) error {
    r.signTime = signTime
    r.Set("sign_time", signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxAuditCreativeAddRequest) GetSignTime() int64 {
    return r.signTime
}
// Creative Setter
// 预审核创意对象
func (r *TaobaoTanxAuditCreativeAddRequest) SetCreative(creative *CreativeParamDTO) error {
    r.creative = creative
    r.Set("creative", creative)
    return nil
}

// Creative Getter
func (r TaobaoTanxAuditCreativeAddRequest) GetCreative() *CreativeParamDTO {
    return r.creative
}
