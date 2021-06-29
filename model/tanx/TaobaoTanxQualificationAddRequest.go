package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交资质接口 API请求
taobao.tanx.qualification.add

dsp客户提交客户资质和行业资质
*/
type TaobaoTanxQualificationAddRequest struct {
    model.Params
    // dsp客户新增资质dto
    qualifications   []Qualification
    // dsp用户memberId
    memberId   int64
    // dsp验证的token
    token   string
    // 签名时间，1970年到现在的秒
    signTime   int64
}

// 初始化TaobaoTanxQualificationAddRequest对象
func NewTaobaoTanxQualificationAddRequest() *TaobaoTanxQualificationAddRequest{
    return &TaobaoTanxQualificationAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTanxQualificationAddRequest) GetApiMethodName() string {
    return "taobao.tanx.qualification.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTanxQualificationAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Qualifications Setter
// dsp客户新增资质dto
func (r *TaobaoTanxQualificationAddRequest) SetQualifications(qualifications []Qualification) error {
    r.qualifications = qualifications
    r.Set("qualifications", qualifications)
    return nil
}

// Qualifications Getter
func (r TaobaoTanxQualificationAddRequest) GetQualifications() []Qualification {
    return r.qualifications
}
// MemberId Setter
// dsp用户memberId
func (r *TaobaoTanxQualificationAddRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxQualificationAddRequest) GetMemberId() int64 {
    return r.memberId
}
// Token Setter
// dsp验证的token
func (r *TaobaoTanxQualificationAddRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r TaobaoTanxQualificationAddRequest) GetToken() string {
    return r.token
}
// SignTime Setter
// 签名时间，1970年到现在的秒
func (r *TaobaoTanxQualificationAddRequest) SetSignTime(signTime int64) error {
    r.signTime = signTime
    r.Set("sign_time", signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxQualificationAddRequest) GetSignTime() int64 {
    return r.signTime
}
