package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest 电信iot自动授权 API请求
// alibaba.ailabs.tmallgenie.third.telecom.autoauth
//
// 电信iot自动授权
type AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest struct {
	model.Params
	// 调用链id,用于全链路跟踪
	_traceId string
	// 电信音箱平台分配给合作方的合作伙伴ID
	_ctPartnerId string
	// 合作方音箱唯一标识，该标识传值为音箱CTEI。
	_deviceId string
	// 访问令牌，用于校验接口调用者身份，做到不可否认性
	_token string
	// 请求时间，格式：yyyy-MM-dd HH:mm:ss
	_time string
	// 电信家居平台用户标识加密字符串
	_tyAccount string
	// 扩展参数，由厂家根据业务自定义内容
	_extraParam string
}

// NewAlibabaAilabsTmallgenieThirdTelecomAutoauthRequest 初始化AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest对象
func NewAlibabaAilabsTmallgenieThirdTelecomAutoauthRequest() *AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest {
	return &AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.third.telecom.autoauth"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceId is TraceId Setter
// 调用链id,用于全链路跟踪
func (r *AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest) SetTraceId(_traceId string) error {
	r._traceId = _traceId
	r.Set("trace_id", _traceId)
	return nil
}

// GetTraceId TraceId Getter
func (r AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest) GetTraceId() string {
	return r._traceId
}

// SetCtPartnerId is CtPartnerId Setter
// 电信音箱平台分配给合作方的合作伙伴ID
func (r *AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest) SetCtPartnerId(_ctPartnerId string) error {
	r._ctPartnerId = _ctPartnerId
	r.Set("ct_partner_id", _ctPartnerId)
	return nil
}

// GetCtPartnerId CtPartnerId Getter
func (r AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest) GetCtPartnerId() string {
	return r._ctPartnerId
}

// SetDeviceId is DeviceId Setter
// 合作方音箱唯一标识，该标识传值为音箱CTEI。
func (r *AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetToken is Token Setter
// 访问令牌，用于校验接口调用者身份，做到不可否认性
func (r *AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest) GetToken() string {
	return r._token
}

// SetTime is Time Setter
// 请求时间，格式：yyyy-MM-dd HH:mm:ss
func (r *AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest) SetTime(_time string) error {
	r._time = _time
	r.Set("time", _time)
	return nil
}

// GetTime Time Getter
func (r AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest) GetTime() string {
	return r._time
}

// SetTyAccount is TyAccount Setter
// 电信家居平台用户标识加密字符串
func (r *AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest) SetTyAccount(_tyAccount string) error {
	r._tyAccount = _tyAccount
	r.Set("ty_account", _tyAccount)
	return nil
}

// GetTyAccount TyAccount Getter
func (r AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest) GetTyAccount() string {
	return r._tyAccount
}

// SetExtraParam is ExtraParam Setter
// 扩展参数，由厂家根据业务自定义内容
func (r *AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest) SetExtraParam(_extraParam string) error {
	r._extraParam = _extraParam
	r.Set("extra_param", _extraParam)
	return nil
}

// GetExtraParam ExtraParam Getter
func (r AlibabaAilabsTmallgenieThirdTelecomAutoauthAPIRequest) GetExtraParam() string {
	return r._extraParam
}
