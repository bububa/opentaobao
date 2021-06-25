package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改资质接口 APIRequest
taobao.tanx.qualification.modify

对dsp上传过的资质进行修改
*/
type TaobaoTanxQualificationModifyRequest struct {
    model.Params

    // dsp客户新增资质dto
    qualifications   []Qualification 

    // dsp用户id
    memberId   int64 

    // dsp用户验证token
    token   string 

    // 1970年到现在的秒
    signTime   int64 

}

func NewTaobaoTanxQualificationModifyRequest() *TaobaoTanxQualificationModifyRequest{
    return &TaobaoTanxQualificationModifyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTanxQualificationModifyRequest) GetApiMethodName() string {
    return "taobao.tanx.qualification.modify"
}

func (r TaobaoTanxQualificationModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTanxQualificationModifyRequest) SetQualifications(qualifications []Qualification) error {
    r.qualifications = qualifications
    r.Set("qualifications", qualifications)
    return nil
}

func (r TaobaoTanxQualificationModifyRequest) GetQualifications() []Qualification {
    return r.qualifications
}

func (r *TaobaoTanxQualificationModifyRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

func (r TaobaoTanxQualificationModifyRequest) GetMemberId() int64 {
    return r.memberId
}

func (r *TaobaoTanxQualificationModifyRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoTanxQualificationModifyRequest) GetToken() string {
    return r.token
}

func (r *TaobaoTanxQualificationModifyRequest) SetSignTime(signTime int64) error {
    r.signTime = signTime
    r.Set("sign_time", signTime)
    return nil
}

func (r TaobaoTanxQualificationModifyRequest) GetSignTime() int64 {
    return r.signTime
}

