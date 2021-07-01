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
type TaobaoTanxQualificationAddAPIRequest struct {
    model.Params
    // dsp客户新增资质dto
    _qualifications   []Qualification
    // dsp用户memberId
    _memberId   int64
    // dsp验证的token
    _token   string
    // 签名时间，1970年到现在的秒
    _signTime   int64
}

// 初始化TaobaoTanxQualificationAddAPIRequest对象
func NewTaobaoTanxQualificationAddRequest() *TaobaoTanxQualificationAddAPIRequest{
    return &TaobaoTanxQualificationAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTanxQualificationAddAPIRequest) GetApiMethodName() string {
    return "taobao.tanx.qualification.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTanxQualificationAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Qualifications Setter
// dsp客户新增资质dto
func (r *TaobaoTanxQualificationAddAPIRequest) SetQualifications(_qualifications []Qualification) error {
    r._qualifications = _qualifications
    r.Set("qualifications", _qualifications)
    return nil
}

// Qualifications Getter
func (r TaobaoTanxQualificationAddAPIRequest) GetQualifications() []Qualification {
    return r._qualifications
}
// MemberId Setter
// dsp用户memberId
func (r *TaobaoTanxQualificationAddAPIRequest) SetMemberId(_memberId int64) error {
    r._memberId = _memberId
    r.Set("member_id", _memberId)
    return nil
}

// MemberId Getter
func (r TaobaoTanxQualificationAddAPIRequest) GetMemberId() int64 {
    return r._memberId
}
// Token Setter
// dsp验证的token
func (r *TaobaoTanxQualificationAddAPIRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoTanxQualificationAddAPIRequest) GetToken() string {
    return r._token
}
// SignTime Setter
// 签名时间，1970年到现在的秒
func (r *TaobaoTanxQualificationAddAPIRequest) SetSignTime(_signTime int64) error {
    r._signTime = _signTime
    r.Set("sign_time", _signTime)
    return nil
}

// SignTime Getter
func (r TaobaoTanxQualificationAddAPIRequest) GetSignTime() int64 {
    return r._signTime
}
