package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTanxQualificationModifyAPIRequest
修改资质接口 API请求
taobao.tanx.qualification.modify

对dsp上传过的资质进行修改 */
type TaobaoTanxQualificationModifyAPIRequest struct {
	model.Params
	// dsp客户新增资质dto
	_qualifications []Qualification
	// dsp用户id
	_memberId int64
	// dsp用户验证token
	_token string
	// 1970年到现在的秒
	_signTime int64
}

// NewTaobaoTanxQualificationModifyRequest 初始化TaobaoTanxQualificationModifyAPIRequest对象
func NewTaobaoTanxQualificationModifyRequest() *TaobaoTanxQualificationModifyAPIRequest {
	return &TaobaoTanxQualificationModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTanxQualificationModifyAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.qualification.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTanxQualificationModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Qualifications Setter
// dsp客户新增资质dto
func (r *TaobaoTanxQualificationModifyAPIRequest) SetQualifications(_qualifications []Qualification) error {
	r._qualifications = _qualifications
	r.Set("qualifications", _qualifications)
	return nil
}

// Get Qualifications Getter
func (r TaobaoTanxQualificationModifyAPIRequest) GetQualifications() []Qualification {
	return r._qualifications
}

// Set is MemberId Setter
// dsp用户id
func (r *TaobaoTanxQualificationModifyAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// Get MemberId Getter
func (r TaobaoTanxQualificationModifyAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// Set is Token Setter
// dsp用户验证token
func (r *TaobaoTanxQualificationModifyAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// Get Token Getter
func (r TaobaoTanxQualificationModifyAPIRequest) GetToken() string {
	return r._token
}

// Set is SignTime Setter
// 1970年到现在的秒
func (r *TaobaoTanxQualificationModifyAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// Get SignTime Getter
func (r TaobaoTanxQualificationModifyAPIRequest) GetSignTime() int64 {
	return r._signTime
}
