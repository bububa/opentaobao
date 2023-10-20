package alisports

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportOauthAlipaygrantAPIRequest 阿里体育会员系统-支付宝授权接口 API请求
// alibaba.alisports.passport.oauth.alipaygrant
//
// 开放给乐心运动使用的支付宝授权接口
type AlibabaAlisportsPassportOauthAlipaygrantAPIRequest struct {
	model.Params
	// 阿里体育分配的用户appkey
	_alispAppKey string
	// 请求接口的时间戳
	_alispTime string
	// 签名字符串
	_alispSign string
	// 调用支付宝登录sdk返回的code
	_authCode string
	// 固定为rich_sports
	_partnerMode string
	// 支付宝的appid
	_appid string
	// 合作方的userid，即用户唯一的id标识
	_appUid string
}

// NewAlibabaAlisportsPassportOauthAlipaygrantRequest 初始化AlibabaAlisportsPassportOauthAlipaygrantAPIRequest对象
func NewAlibabaAlisportsPassportOauthAlipaygrantRequest() *AlibabaAlisportsPassportOauthAlipaygrantAPIRequest {
	return &AlibabaAlisportsPassportOauthAlipaygrantAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlisportsPassportOauthAlipaygrantAPIRequest) Reset() {
	r._alispAppKey = ""
	r._alispTime = ""
	r._alispSign = ""
	r._authCode = ""
	r._partnerMode = ""
	r._appid = ""
	r._appUid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportOauthAlipaygrantAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.passport.oauth.alipaygrant"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportOauthAlipaygrantAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlisportsPassportOauthAlipaygrantAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlispAppKey is AlispAppKey Setter
// 阿里体育分配的用户appkey
func (r *AlibabaAlisportsPassportOauthAlipaygrantAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// GetAlispAppKey AlispAppKey Getter
func (r AlibabaAlisportsPassportOauthAlipaygrantAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// SetAlispTime is AlispTime Setter
// 请求接口的时间戳
func (r *AlibabaAlisportsPassportOauthAlipaygrantAPIRequest) SetAlispTime(_alispTime string) error {
	r._alispTime = _alispTime
	r.Set("alisp_time", _alispTime)
	return nil
}

// GetAlispTime AlispTime Getter
func (r AlibabaAlisportsPassportOauthAlipaygrantAPIRequest) GetAlispTime() string {
	return r._alispTime
}

// SetAlispSign is AlispSign Setter
// 签名字符串
func (r *AlibabaAlisportsPassportOauthAlipaygrantAPIRequest) SetAlispSign(_alispSign string) error {
	r._alispSign = _alispSign
	r.Set("alisp_sign", _alispSign)
	return nil
}

// GetAlispSign AlispSign Getter
func (r AlibabaAlisportsPassportOauthAlipaygrantAPIRequest) GetAlispSign() string {
	return r._alispSign
}

// SetAuthCode is AuthCode Setter
// 调用支付宝登录sdk返回的code
func (r *AlibabaAlisportsPassportOauthAlipaygrantAPIRequest) SetAuthCode(_authCode string) error {
	r._authCode = _authCode
	r.Set("auth_code", _authCode)
	return nil
}

// GetAuthCode AuthCode Getter
func (r AlibabaAlisportsPassportOauthAlipaygrantAPIRequest) GetAuthCode() string {
	return r._authCode
}

// SetPartnerMode is PartnerMode Setter
// 固定为rich_sports
func (r *AlibabaAlisportsPassportOauthAlipaygrantAPIRequest) SetPartnerMode(_partnerMode string) error {
	r._partnerMode = _partnerMode
	r.Set("partner_mode", _partnerMode)
	return nil
}

// GetPartnerMode PartnerMode Getter
func (r AlibabaAlisportsPassportOauthAlipaygrantAPIRequest) GetPartnerMode() string {
	return r._partnerMode
}

// SetAppid is Appid Setter
// 支付宝的appid
func (r *AlibabaAlisportsPassportOauthAlipaygrantAPIRequest) SetAppid(_appid string) error {
	r._appid = _appid
	r.Set("appid", _appid)
	return nil
}

// GetAppid Appid Getter
func (r AlibabaAlisportsPassportOauthAlipaygrantAPIRequest) GetAppid() string {
	return r._appid
}

// SetAppUid is AppUid Setter
// 合作方的userid，即用户唯一的id标识
func (r *AlibabaAlisportsPassportOauthAlipaygrantAPIRequest) SetAppUid(_appUid string) error {
	r._appUid = _appUid
	r.Set("app_uid", _appUid)
	return nil
}

// GetAppUid AppUid Getter
func (r AlibabaAlisportsPassportOauthAlipaygrantAPIRequest) GetAppUid() string {
	return r._appUid
}

var poolAlibabaAlisportsPassportOauthAlipaygrantAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlisportsPassportOauthAlipaygrantRequest()
	},
}

// GetAlibabaAlisportsPassportOauthAlipaygrantRequest 从 sync.Pool 获取 AlibabaAlisportsPassportOauthAlipaygrantAPIRequest
func GetAlibabaAlisportsPassportOauthAlipaygrantAPIRequest() *AlibabaAlisportsPassportOauthAlipaygrantAPIRequest {
	return poolAlibabaAlisportsPassportOauthAlipaygrantAPIRequest.Get().(*AlibabaAlisportsPassportOauthAlipaygrantAPIRequest)
}

// ReleaseAlibabaAlisportsPassportOauthAlipaygrantAPIRequest 将 AlibabaAlisportsPassportOauthAlipaygrantAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlisportsPassportOauthAlipaygrantAPIRequest(v *AlibabaAlisportsPassportOauthAlipaygrantAPIRequest) {
	v.Reset()
	poolAlibabaAlisportsPassportOauthAlipaygrantAPIRequest.Put(v)
}
