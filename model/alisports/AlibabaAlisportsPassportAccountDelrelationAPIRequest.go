package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportAccountDelrelationAPIRequest 阿里体育会员系统--取消三方关联接口 API请求
// alibaba.alisports.passport.account.delrelation
//
// 阿里体育会员系统--取消三方关联接口
type AlibabaAlisportsPassportAccountDelrelationAPIRequest struct {
	model.Params
	// 业务appkey
	_alispAppKey string
	// 调用时间戳
	_alispTime string
	// 签名字符串
	_alispSign string
	// 合作方用户ID
	_appUid string
	// 阿里体育会员id
	_aliuid string
}

// NewAlibabaAlisportsPassportAccountDelrelationRequest 初始化AlibabaAlisportsPassportAccountDelrelationAPIRequest对象
func NewAlibabaAlisportsPassportAccountDelrelationRequest() *AlibabaAlisportsPassportAccountDelrelationAPIRequest {
	return &AlibabaAlisportsPassportAccountDelrelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportAccountDelrelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.passport.account.delrelation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportAccountDelrelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlisportsPassportAccountDelrelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlispAppKey is AlispAppKey Setter
// 业务appkey
func (r *AlibabaAlisportsPassportAccountDelrelationAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// GetAlispAppKey AlispAppKey Getter
func (r AlibabaAlisportsPassportAccountDelrelationAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// SetAlispTime is AlispTime Setter
// 调用时间戳
func (r *AlibabaAlisportsPassportAccountDelrelationAPIRequest) SetAlispTime(_alispTime string) error {
	r._alispTime = _alispTime
	r.Set("alisp_time", _alispTime)
	return nil
}

// GetAlispTime AlispTime Getter
func (r AlibabaAlisportsPassportAccountDelrelationAPIRequest) GetAlispTime() string {
	return r._alispTime
}

// SetAlispSign is AlispSign Setter
// 签名字符串
func (r *AlibabaAlisportsPassportAccountDelrelationAPIRequest) SetAlispSign(_alispSign string) error {
	r._alispSign = _alispSign
	r.Set("alisp_sign", _alispSign)
	return nil
}

// GetAlispSign AlispSign Getter
func (r AlibabaAlisportsPassportAccountDelrelationAPIRequest) GetAlispSign() string {
	return r._alispSign
}

// SetAppUid is AppUid Setter
// 合作方用户ID
func (r *AlibabaAlisportsPassportAccountDelrelationAPIRequest) SetAppUid(_appUid string) error {
	r._appUid = _appUid
	r.Set("app_uid", _appUid)
	return nil
}

// GetAppUid AppUid Getter
func (r AlibabaAlisportsPassportAccountDelrelationAPIRequest) GetAppUid() string {
	return r._appUid
}

// SetAliuid is Aliuid Setter
// 阿里体育会员id
func (r *AlibabaAlisportsPassportAccountDelrelationAPIRequest) SetAliuid(_aliuid string) error {
	r._aliuid = _aliuid
	r.Set("aliuid", _aliuid)
	return nil
}

// GetAliuid Aliuid Getter
func (r AlibabaAlisportsPassportAccountDelrelationAPIRequest) GetAliuid() string {
	return r._aliuid
}
