package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交资质接口 APIRequest
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

func NewTaobaoTanxQualificationAddRequest() *TaobaoTanxQualificationAddRequest{
    return &TaobaoTanxQualificationAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTanxQualificationAddRequest) GetApiMethodName() string {
    return "taobao.tanx.qualification.add"
}

func (r TaobaoTanxQualificationAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTanxQualificationAddRequest) SetQualifications(qualifications []Qualification) error {
    r.qualifications = qualifications
    r.Set("qualifications", qualifications)
    return nil
}

func (r TaobaoTanxQualificationAddRequest) GetQualifications() []Qualification {
    return r.qualifications
}

func (r *TaobaoTanxQualificationAddRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

func (r TaobaoTanxQualificationAddRequest) GetMemberId() int64 {
    return r.memberId
}

func (r *TaobaoTanxQualificationAddRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoTanxQualificationAddRequest) GetToken() string {
    return r.token
}

func (r *TaobaoTanxQualificationAddRequest) SetSignTime(signTime int64) error {
    r.signTime = signTime
    r.Set("sign_time", signTime)
    return nil
}

func (r TaobaoTanxQualificationAddRequest) GetSignTime() int64 {
    return r.signTime
}

