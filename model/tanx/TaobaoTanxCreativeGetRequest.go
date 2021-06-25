package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单个审核创意状态 APIRequest
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

func NewTaobaoTanxCreativeGetRequest() *TaobaoTanxCreativeGetRequest{
    return &TaobaoTanxCreativeGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTanxCreativeGetRequest) GetApiMethodName() string {
    return "taobao.tanx.creative.get"
}

func (r TaobaoTanxCreativeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTanxCreativeGetRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

func (r TaobaoTanxCreativeGetRequest) GetMemberId() int64 {
    return r.memberId
}

func (r *TaobaoTanxCreativeGetRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoTanxCreativeGetRequest) GetToken() string {
    return r.token
}

func (r *TaobaoTanxCreativeGetRequest) SetSignTime(signTime int64) error {
    r.signTime = signTime
    r.Set("sign_time", signTime)
    return nil
}

func (r TaobaoTanxCreativeGetRequest) GetSignTime() int64 {
    return r.signTime
}

func (r *TaobaoTanxCreativeGetRequest) SetCreativeId(creativeId string) error {
    r.creativeId = creativeId
    r.Set("creative_id", creativeId)
    return nil
}

func (r TaobaoTanxCreativeGetRequest) GetCreativeId() string {
    return r.creativeId
}

