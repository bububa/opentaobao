package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创意预审新增接口 APIRequest
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

func NewTaobaoTanxAuditCreativeAddRequest() *TaobaoTanxAuditCreativeAddRequest{
    return &TaobaoTanxAuditCreativeAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTanxAuditCreativeAddRequest) GetApiMethodName() string {
    return "taobao.tanx.audit.creative.add"
}

func (r TaobaoTanxAuditCreativeAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTanxAuditCreativeAddRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

func (r TaobaoTanxAuditCreativeAddRequest) GetMemberId() int64 {
    return r.memberId
}

func (r *TaobaoTanxAuditCreativeAddRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoTanxAuditCreativeAddRequest) GetToken() string {
    return r.token
}

func (r *TaobaoTanxAuditCreativeAddRequest) SetSignTime(signTime int64) error {
    r.signTime = signTime
    r.Set("sign_time", signTime)
    return nil
}

func (r TaobaoTanxAuditCreativeAddRequest) GetSignTime() int64 {
    return r.signTime
}

func (r *TaobaoTanxAuditCreativeAddRequest) SetCreative(creative *CreativeParamDTO) error {
    r.creative = creative
    r.Set("creative", creative)
    return nil
}

func (r TaobaoTanxAuditCreativeAddRequest) GetCreative() *CreativeParamDTO {
    return r.creative
}

