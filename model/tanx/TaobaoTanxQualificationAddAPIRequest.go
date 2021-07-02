package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTanxQualificationAddAPIRequest 提交资质接口 API请求
// taobao.tanx.qualification.add
//
// dsp客户提交客户资质和行业资质
type TaobaoTanxQualificationAddAPIRequest struct {
	model.Params
	// dsp客户新增资质dto
	_qualifications []Qualification
	// dsp用户memberId
	_memberId int64
	// dsp验证的token
	_token string
	// 签名时间，1970年到现在的秒
	_signTime int64
}

// NewTaobaoTanxQualificationAddRequest 初始化TaobaoTanxQualificationAddAPIRequest对象
func NewTaobaoTanxQualificationAddRequest() *TaobaoTanxQualificationAddAPIRequest {
	return &TaobaoTanxQualificationAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTanxQualificationAddAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.qualification.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTanxQualificationAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetQualifications is Qualifications Setter
// dsp客户新增资质dto
func (r *TaobaoTanxQualificationAddAPIRequest) SetQualifications(_qualifications []Qualification) error {
	r._qualifications = _qualifications
	r.Set("qualifications", _qualifications)
	return nil
}

// GetQualifications Qualifications Getter
func (r TaobaoTanxQualificationAddAPIRequest) GetQualifications() []Qualification {
	return r._qualifications
}

// SetMemberId is MemberId Setter
// dsp用户memberId
func (r *TaobaoTanxQualificationAddAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r TaobaoTanxQualificationAddAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// SetToken is Token Setter
// dsp验证的token
func (r *TaobaoTanxQualificationAddAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaoTanxQualificationAddAPIRequest) GetToken() string {
	return r._token
}

// SetSignTime is SignTime Setter
// 签名时间，1970年到现在的秒
func (r *TaobaoTanxQualificationAddAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TaobaoTanxQualificationAddAPIRequest) GetSignTime() int64 {
	return r._signTime
}
