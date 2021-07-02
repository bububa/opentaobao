package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyMiniappUserGetAPIRequest 零售云小程序获取登录用户信息 API请求
// alibaba.lsy.miniapp.user.get
//
// 零售云小程序，通过授权码获取登录的卖家账号信息
type AlibabaLsyMiniappUserGetAPIRequest struct {
	model.Params
	// 当前时间戳，毫秒
	_timeStamp string
	// 获取用户信息的授权码，在小程序中获取
	_code string
	// 请求参数签名，sha1(所有入参+appSecret，按字符串升序排列)
	_signature string
	// 系统分配的小程序ID
	_appId string
}

// NewAlibabaLsyMiniappUserGetRequest 初始化AlibabaLsyMiniappUserGetAPIRequest对象
func NewAlibabaLsyMiniappUserGetRequest() *AlibabaLsyMiniappUserGetAPIRequest {
	return &AlibabaLsyMiniappUserGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLsyMiniappUserGetAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.miniapp.user.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLsyMiniappUserGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTimeStamp is TimeStamp Setter
// 当前时间戳，毫秒
func (r *AlibabaLsyMiniappUserGetAPIRequest) SetTimeStamp(_timeStamp string) error {
	r._timeStamp = _timeStamp
	r.Set("time_stamp", _timeStamp)
	return nil
}

// GetTimeStamp TimeStamp Getter
func (r AlibabaLsyMiniappUserGetAPIRequest) GetTimeStamp() string {
	return r._timeStamp
}

// SetCode is Code Setter
// 获取用户信息的授权码，在小程序中获取
func (r *AlibabaLsyMiniappUserGetAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaLsyMiniappUserGetAPIRequest) GetCode() string {
	return r._code
}

// SetSignature is Signature Setter
// 请求参数签名，sha1(所有入参+appSecret，按字符串升序排列)
func (r *AlibabaLsyMiniappUserGetAPIRequest) SetSignature(_signature string) error {
	r._signature = _signature
	r.Set("signature", _signature)
	return nil
}

// GetSignature Signature Getter
func (r AlibabaLsyMiniappUserGetAPIRequest) GetSignature() string {
	return r._signature
}

// SetAppId is AppId Setter
// 系统分配的小程序ID
func (r *AlibabaLsyMiniappUserGetAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r AlibabaLsyMiniappUserGetAPIRequest) GetAppId() string {
	return r._appId
}
