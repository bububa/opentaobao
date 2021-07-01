package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsPassportAccountBindthirdidAPIRequest
阿里体育三方ID绑定接口 API请求
alibaba.alisports.passport.account.bindthirdid

阿里体育三方ID绑定接口 */
type AlibabaAlisportsPassportAccountBindthirdidAPIRequest struct {
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

// NewAlibabaAlisportsPassportAccountBindthirdidRequest 初始化AlibabaAlisportsPassportAccountBindthirdidAPIRequest对象
func NewAlibabaAlisportsPassportAccountBindthirdidRequest() *AlibabaAlisportsPassportAccountBindthirdidAPIRequest {
	return &AlibabaAlisportsPassportAccountBindthirdidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportAccountBindthirdidAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.passport.account.bindthirdid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportAccountBindthirdidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AlispAppKey Setter
// 业务方appkey
func (r *AlibabaAlisportsPassportAccountBindthirdidAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// Get AlispAppKey Getter
func (r AlibabaAlisportsPassportAccountBindthirdidAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// Set is AlispTime Setter
// 时间戳精确到秒
func (r *AlibabaAlisportsPassportAccountBindthirdidAPIRequest) SetAlispTime(_alispTime string) error {
	r._alispTime = _alispTime
	r.Set("alisp_time", _alispTime)
	return nil
}

// Get AlispTime Getter
func (r AlibabaAlisportsPassportAccountBindthirdidAPIRequest) GetAlispTime() string {
	return r._alispTime
}

// Set is AlispSign Setter
// 接口签名
func (r *AlibabaAlisportsPassportAccountBindthirdidAPIRequest) SetAlispSign(_alispSign string) error {
	r._alispSign = _alispSign
	r.Set("alisp_sign", _alispSign)
	return nil
}

// Get AlispSign Getter
func (r AlibabaAlisportsPassportAccountBindthirdidAPIRequest) GetAlispSign() string {
	return r._alispSign
}

// Set is Aliuid Setter
// 阿里体育用户ID
func (r *AlibabaAlisportsPassportAccountBindthirdidAPIRequest) SetAliuid(_aliuid string) error {
	r._aliuid = _aliuid
	r.Set("aliuid", _aliuid)
	return nil
}

// Get Aliuid Getter
func (r AlibabaAlisportsPassportAccountBindthirdidAPIRequest) GetAliuid() string {
	return r._aliuid
}

// Set is AppUid Setter
// 三方ID
func (r *AlibabaAlisportsPassportAccountBindthirdidAPIRequest) SetAppUid(_appUid string) error {
	r._appUid = _appUid
	r.Set("app_uid", _appUid)
	return nil
}

// Get AppUid Getter
func (r AlibabaAlisportsPassportAccountBindthirdidAPIRequest) GetAppUid() string {
	return r._appUid
}

// Set is Mobile Setter
// 手机号
func (r *AlibabaAlisportsPassportAccountBindthirdidAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// Get Mobile Getter
func (r AlibabaAlisportsPassportAccountBindthirdidAPIRequest) GetMobile() string {
	return r._mobile
}
