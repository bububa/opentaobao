package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalisportspassportaccountbindthirdidAPIRequest 阿里体育三方ID绑定接口 API请求
// alibaba.alisports.passport.account.bindthirdid
//
// 阿里体育三方ID绑定接口
type AlibabaalisportspassportaccountbindthirdidAPIRequest struct {
	model.Params
	// 业务方appkey
	_alispAppKey string
	// 时间戳精确到秒
	_alispTime string
	// 接口签名
	_alispSign string
	// 阿里体育用户ID
	_aliuid string
	// 三方ID
	_appUid string
	// 手机号
	_mobile string
}

// NewAlibabaalisportspassportaccountbindthirdidRequest 初始化AlibabaalisportspassportaccountbindthirdidAPIRequest对象
func NewAlibabaalisportspassportaccountbindthirdidRequest() *AlibabaalisportspassportaccountbindthirdidAPIRequest {
	return &AlibabaalisportspassportaccountbindthirdidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalisportspassportaccountbindthirdidAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.passport.account.bindthirdid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalisportspassportaccountbindthirdidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalisportspassportaccountbindthirdidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlispAppKey is AlispAppKey Setter
// 业务方appkey
func (r *AlibabaalisportspassportaccountbindthirdidAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// GetAlispAppKey AlispAppKey Getter
func (r AlibabaalisportspassportaccountbindthirdidAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// SetAlispTime is AlispTime Setter
// 时间戳精确到秒
func (r *AlibabaalisportspassportaccountbindthirdidAPIRequest) SetAlispTime(_alispTime string) error {
	r._alispTime = _alispTime
	r.Set("alisp_time", _alispTime)
	return nil
}

// GetAlispTime AlispTime Getter
func (r AlibabaalisportspassportaccountbindthirdidAPIRequest) GetAlispTime() string {
	return r._alispTime
}

// SetAlispSign is AlispSign Setter
// 接口签名
func (r *AlibabaalisportspassportaccountbindthirdidAPIRequest) SetAlispSign(_alispSign string) error {
	r._alispSign = _alispSign
	r.Set("alisp_sign", _alispSign)
	return nil
}

// GetAlispSign AlispSign Getter
func (r AlibabaalisportspassportaccountbindthirdidAPIRequest) GetAlispSign() string {
	return r._alispSign
}

// SetAliuid is Aliuid Setter
// 阿里体育用户ID
func (r *AlibabaalisportspassportaccountbindthirdidAPIRequest) SetAliuid(_aliuid string) error {
	r._aliuid = _aliuid
	r.Set("aliuid", _aliuid)
	return nil
}

// GetAliuid Aliuid Getter
func (r AlibabaalisportspassportaccountbindthirdidAPIRequest) GetAliuid() string {
	return r._aliuid
}

// SetAppUid is AppUid Setter
// 三方ID
func (r *AlibabaalisportspassportaccountbindthirdidAPIRequest) SetAppUid(_appUid string) error {
	r._appUid = _appUid
	r.Set("app_uid", _appUid)
	return nil
}

// GetAppUid AppUid Getter
func (r AlibabaalisportspassportaccountbindthirdidAPIRequest) GetAppUid() string {
	return r._appUid
}

// SetMobile is Mobile Setter
// 手机号
func (r *AlibabaalisportspassportaccountbindthirdidAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r AlibabaalisportspassportaccountbindthirdidAPIRequest) GetMobile() string {
	return r._mobile
}
