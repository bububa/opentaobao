package vaccin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinNoticeUserBindAPIRequest 支付宝疫苗绑定接种人 API请求
// alibaba.health.vaccin.notice.user.bind
//
// 支付宝疫苗绑定接种人
type AlibabaHealthVaccinNoticeUserBindAPIRequest struct {
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

// NewAlibabaHealthVaccinNoticeUserBindRequest 初始化AlibabaHealthVaccinNoticeUserBindAPIRequest对象
func NewAlibabaHealthVaccinNoticeUserBindRequest() *AlibabaHealthVaccinNoticeUserBindAPIRequest {
	return &AlibabaHealthVaccinNoticeUserBindAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHealthVaccinNoticeUserBindAPIRequest) Reset() {
	r._bindUsers = r._bindUsers[:0]
	r._alipayUserId = ""
	r._outerUserId = ""
	r._mobile = ""
	r._appChannel = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinNoticeUserBindAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.notice.user.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinNoticeUserBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHealthVaccinNoticeUserBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBindUsers is BindUsers Setter
// 绑定人信息list
func (r *AlibabaHealthVaccinNoticeUserBindAPIRequest) SetBindUsers(_bindUsers []AlipayVaccineUserBindDto) error {
	r._bindUsers = _bindUsers
	r.Set("bind_users", _bindUsers)
	return nil
}

// GetBindUsers BindUsers Getter
func (r AlibabaHealthVaccinNoticeUserBindAPIRequest) GetBindUsers() []AlipayVaccineUserBindDto {
	return r._bindUsers
}

// SetAlipayUserId is AlipayUserId Setter
// 支付宝ID
func (r *AlibabaHealthVaccinNoticeUserBindAPIRequest) SetAlipayUserId(_alipayUserId string) error {
	r._alipayUserId = _alipayUserId
	r.Set("alipay_user_id", _alipayUserId)
	return nil
}

// GetAlipayUserId AlipayUserId Getter
func (r AlibabaHealthVaccinNoticeUserBindAPIRequest) GetAlipayUserId() string {
	return r._alipayUserId
}

// SetOuterUserId is OuterUserId Setter
// ISV 侧用户 ID
func (r *AlibabaHealthVaccinNoticeUserBindAPIRequest) SetOuterUserId(_outerUserId string) error {
	r._outerUserId = _outerUserId
	r.Set("outer_user_id", _outerUserId)
	return nil
}

// GetOuterUserId OuterUserId Getter
func (r AlibabaHealthVaccinNoticeUserBindAPIRequest) GetOuterUserId() string {
	return r._outerUserId
}

// SetMobile is Mobile Setter
// 联系电话
func (r *AlibabaHealthVaccinNoticeUserBindAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r AlibabaHealthVaccinNoticeUserBindAPIRequest) GetMobile() string {
	return r._mobile
}

// SetAppChannel is AppChannel Setter
// 用户来源：alipay或yl
func (r *AlibabaHealthVaccinNoticeUserBindAPIRequest) SetAppChannel(_appChannel string) error {
	r._appChannel = _appChannel
	r.Set("app_channel", _appChannel)
	return nil
}

// GetAppChannel AppChannel Getter
func (r AlibabaHealthVaccinNoticeUserBindAPIRequest) GetAppChannel() string {
	return r._appChannel
}

var poolAlibabaHealthVaccinNoticeUserBindAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHealthVaccinNoticeUserBindRequest()
	},
}

// GetAlibabaHealthVaccinNoticeUserBindRequest 从 sync.Pool 获取 AlibabaHealthVaccinNoticeUserBindAPIRequest
func GetAlibabaHealthVaccinNoticeUserBindAPIRequest() *AlibabaHealthVaccinNoticeUserBindAPIRequest {
	return poolAlibabaHealthVaccinNoticeUserBindAPIRequest.Get().(*AlibabaHealthVaccinNoticeUserBindAPIRequest)
}

// ReleaseAlibabaHealthVaccinNoticeUserBindAPIRequest 将 AlibabaHealthVaccinNoticeUserBindAPIRequest 放入 sync.Pool
func ReleaseAlibabaHealthVaccinNoticeUserBindAPIRequest(v *AlibabaHealthVaccinNoticeUserBindAPIRequest) {
	v.Reset()
	poolAlibabaHealthVaccinNoticeUserBindAPIRequest.Put(v)
}
