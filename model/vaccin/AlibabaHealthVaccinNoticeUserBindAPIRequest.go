package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthVaccinNoticeUserBindAPIRequest
支付宝疫苗绑定接种人 API请求
alibaba.health.vaccin.notice.user.bind

支付宝疫苗绑定接种人 */
type AlibabaHealthVaccinNoticeUserBindAPIRequest struct {
	model.Params
	// 支付宝ID
	_alipayUserId string
	// 绑定人信息list
	_bindUsers []AlipayVaccineUserBindDto
	// ISV 侧用户 ID
	_outerUserId string
	// 联系电话
	_mobile string
	// 用户来源：alipay或yl
	_appChannel string
}

// NewAlibabaHealthVaccinNoticeUserBindRequest 初始化AlibabaHealthVaccinNoticeUserBindAPIRequest对象
func NewAlibabaHealthVaccinNoticeUserBindRequest() *AlibabaHealthVaccinNoticeUserBindAPIRequest {
	return &AlibabaHealthVaccinNoticeUserBindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeUserBindAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.notice.user.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeUserBindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AlipayUserId Setter
// 支付宝ID
func (r *AlibabaHealthVaccinNoticeUserBindAPIRequest) SetAlipayUserId(_alipayUserId string) error {
	r._alipayUserId = _alipayUserId
	r.Set("alipay_user_id", _alipayUserId)
	return nil
}

// Get AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeUserBindAPIRequest) GetAlipayUserId() string {
	return r._alipayUserId
}

// Set is BindUsers Setter
// 绑定人信息list
func (r *AlibabaHealthVaccinNoticeUserBindAPIRequest) SetBindUsers(_bindUsers []AlipayVaccineUserBindDto) error {
	r._bindUsers = _bindUsers
	r.Set("bind_users", _bindUsers)
	return nil
}

// Get BindUsers Getter
func (r AlibabaHealthVaccinNoticeUserBindAPIRequest) GetBindUsers() []AlipayVaccineUserBindDto {
	return r._bindUsers
}

// Set is OuterUserId Setter
// ISV 侧用户 ID
func (r *AlibabaHealthVaccinNoticeUserBindAPIRequest) SetOuterUserId(_outerUserId string) error {
	r._outerUserId = _outerUserId
	r.Set("outer_user_id", _outerUserId)
	return nil
}

// Get OuterUserId Getter
func (r AlibabaHealthVaccinNoticeUserBindAPIRequest) GetOuterUserId() string {
	return r._outerUserId
}

// Set is Mobile Setter
// 联系电话
func (r *AlibabaHealthVaccinNoticeUserBindAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// Get Mobile Getter
func (r AlibabaHealthVaccinNoticeUserBindAPIRequest) GetMobile() string {
	return r._mobile
}

// Set is AppChannel Setter
// 用户来源：alipay或yl
func (r *AlibabaHealthVaccinNoticeUserBindAPIRequest) SetAppChannel(_appChannel string) error {
	r._appChannel = _appChannel
	r.Set("app_channel", _appChannel)
	return nil
}

// Get AppChannel Getter
func (r AlibabaHealthVaccinNoticeUserBindAPIRequest) GetAppChannel() string {
	return r._appChannel
}
