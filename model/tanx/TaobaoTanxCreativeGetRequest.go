package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单个审核创意状态 API请求
taobao.tanx.creative.get

获取单个审核创意状态
*/
type TaobaoTanxCreativeGetRequest struct {
    model.Params
    // DSP的memberId
    memberId   int64
    // dsp用户身份认证的TOKEN
    token   string
    // 当前时间戳，1970-01-01后的秒数
    signTime   int64
    // 创意ID
    creativeId   string
}

// 初始化TaobaoTanxCreativeGetRequest对象
func NewTaobaoTanxCreativeGetRequest() *TaobaoTanxCreativeGetRequest{
    return &TaobaoTanxCreativeGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTanxCreativeGetRequest) GetApiMethodName() string {
    return "taobao.tanx.creative.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTanxCreativeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MemberId Setter
// DSP的memberId
func (r *TaobaoTanxCreativeGetRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxCreativeGetRequest) GetMemberId() int64 {
    return r.memberId
}
// Token Setter
// dsp用户身份认证的TOKEN
func (r *TaobaoTanxCreativeGetRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r TaobaoTanxCreativeGetRequest) GetToken() string {
    return r.token
}
// SignTime Setter
// 当前时间戳，1970-01-01后的秒数
func (r *TaobaoTanxCreativeGetRequest) SetSignTime(signTime int64) error {
    r.signTime = signTime
    r.Set("sign_time", signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxCreativeGetRequest) GetSignTime() int64 {
    return r.signTime
}
// CreativeId Setter
// 创意ID
func (r *TaobaoTanxCreativeGetRequest) SetCreativeId(creativeId string) error {
    r.creativeId = creativeId
    r.Set("creative_id", creativeId)
    return nil
}

// CreativeId Getter
func (r TaobaoTanxCreativeGetRequest) GetCreativeId() string {
    return r.creativeId
}
