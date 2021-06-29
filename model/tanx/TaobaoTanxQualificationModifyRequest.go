package tanx

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改资质接口 API请求
taobao.tanx.qualification.modify

对dsp上传过的资质进行修改
*/
type TaobaoTanxQualificationModifyRequest struct {
    model.Params
    // dsp客户新增资质dto
    _qualifications   []Qualification
    // dsp用户id
    _memberId   int64
    // dsp用户验证token
    _token   string
    // 1970年到现在的秒
    _signTime   int64
}

// 初始化TaobaoTanxQualificationModifyRequest对象
func NewTaobaoTanxQualificationModifyRequest() *TaobaoTanxQualificationModifyRequest{
    return &TaobaoTanxQualificationModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTanxQualificationModifyRequest) GetApiMethodName() string {
    return "taobao.tanx.qualification.modify"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTanxQualificationModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Qualifications Setter
// dsp客户新增资质dto
func (r *TaobaoTanxQualificationModifyRequest) SetQualifications(_qualifications []Qualification) error {
    r._qualifications = _qualifications
    r.Set("qualifications", _qualifications)
    return nil
}

// Qualifications Getter
func (r TaobaoTanxQualificationModifyRequest) GetQualifications() []Qualification {
    return r._qualifications
}
// MemberId Setter
// dsp用户id
func (r *TaobaoTanxQualificationModifyRequest) SetMemberId(_memberId int64) error {
    r._memberId = _memberId
    r.Set("member_id", _memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxQualificationModifyRequest) GetMemberId() int64 {
    return r._memberId
}
// Token Setter
// dsp用户验证token
func (r *TaobaoTanxQualificationModifyRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoTanxQualificationModifyRequest) GetToken() string {
    return r._token
}
// SignTime Setter
// 1970年到现在的秒
func (r *TaobaoTanxQualificationModifyRequest) SetSignTime(_signTime int64) error {
    r._signTime = _signTime
    r.Set("sign_time", _signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxQualificationModifyRequest) GetSignTime() int64 {
    return r._signTime
}
