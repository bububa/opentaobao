package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahealthvaccinnoticeuserbindAPIRequest 支付宝疫苗绑定接种人 API请求
// alibaba.health.vaccin.notice.user.bind
//
// 支付宝疫苗绑定接种人
type AlibabahealthvaccinnoticeuserbindAPIRequest struct {
	model.Params
	// 绑定人信息list
	_bindUsers []AlipayVaccineUserBindDto
	// 支付宝ID
	_alipayUserId string
	// ISV 侧用户 ID
	_outerUserId string
	// 联系电话
	_mobile string
	// 用户来源：alipay或yl
	_appChannel string
}

// NewAlibabahealthvaccinnoticeuserbindRequest 初始化AlibabahealthvaccinnoticeuserbindAPIRequest对象
func NewAlibabahealthvaccinnoticeuserbindRequest() *AlibabahealthvaccinnoticeuserbindAPIRequest {
	return &AlibabahealthvaccinnoticeuserbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahealthvaccinnoticeuserbindAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.notice.user.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahealthvaccinnoticeuserbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahealthvaccinnoticeuserbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBindUsers is BindUsers Setter
// 绑定人信息list
func (r *AlibabahealthvaccinnoticeuserbindAPIRequest) SetBindUsers(_bindUsers []AlipayVaccineUserBindDto) error {
	r._bindUsers = _bindUsers
	r.Set("bind_users", _bindUsers)
	return nil
}

// GetBindUsers BindUsers Getter
func (r AlibabahealthvaccinnoticeuserbindAPIRequest) GetBindUsers() []AlipayVaccineUserBindDto {
	return r._bindUsers
}

// SetAlipayUserId is AlipayUserId Setter
// 支付宝ID
func (r *AlibabahealthvaccinnoticeuserbindAPIRequest) SetAlipayUserId(_alipayUserId string) error {
	r._alipayUserId = _alipayUserId
	r.Set("alipay_user_id", _alipayUserId)
	return nil
}

// GetAlipayUserId AlipayUserId Getter
func (r AlibabahealthvaccinnoticeuserbindAPIRequest) GetAlipayUserId() string {
	return r._alipayUserId
}

// SetOuterUserId is OuterUserId Setter
// ISV 侧用户 ID
func (r *AlibabahealthvaccinnoticeuserbindAPIRequest) SetOuterUserId(_outerUserId string) error {
	r._outerUserId = _outerUserId
	r.Set("outer_user_id", _outerUserId)
	return nil
}

// GetOuterUserId OuterUserId Getter
func (r AlibabahealthvaccinnoticeuserbindAPIRequest) GetOuterUserId() string {
	return r._outerUserId
}

// SetMobile is Mobile Setter
// 联系电话
func (r *AlibabahealthvaccinnoticeuserbindAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r AlibabahealthvaccinnoticeuserbindAPIRequest) GetMobile() string {
	return r._mobile
}

// SetAppChannel is AppChannel Setter
// 用户来源：alipay或yl
func (r *AlibabahealthvaccinnoticeuserbindAPIRequest) SetAppChannel(_appChannel string) error {
	r._appChannel = _appChannel
	r.Set("app_channel", _appChannel)
	return nil
}

// GetAppChannel AppChannel Getter
func (r AlibabahealthvaccinnoticeuserbindAPIRequest) GetAppChannel() string {
	return r._appChannel
}
