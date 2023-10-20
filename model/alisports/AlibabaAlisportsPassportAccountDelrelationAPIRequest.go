package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalisportspassportaccountdelrelationAPIRequest 阿里体育会员系统--取消三方关联接口 API请求
// alibaba.alisports.passport.account.delrelation
//
// 阿里体育会员系统--取消三方关联接口
type AlibabaalisportspassportaccountdelrelationAPIRequest struct {
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

// NewAlibabaalisportspassportaccountdelrelationRequest 初始化AlibabaalisportspassportaccountdelrelationAPIRequest对象
func NewAlibabaalisportspassportaccountdelrelationRequest() *AlibabaalisportspassportaccountdelrelationAPIRequest {
	return &AlibabaalisportspassportaccountdelrelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalisportspassportaccountdelrelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.passport.account.delrelation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalisportspassportaccountdelrelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalisportspassportaccountdelrelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlispAppKey is AlispAppKey Setter
// 业务appkey
func (r *AlibabaalisportspassportaccountdelrelationAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// GetAlispAppKey AlispAppKey Getter
func (r AlibabaalisportspassportaccountdelrelationAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// SetAlispTime is AlispTime Setter
// 调用时间戳
func (r *AlibabaalisportspassportaccountdelrelationAPIRequest) SetAlispTime(_alispTime string) error {
	r._alispTime = _alispTime
	r.Set("alisp_time", _alispTime)
	return nil
}

// GetAlispTime AlispTime Getter
func (r AlibabaalisportspassportaccountdelrelationAPIRequest) GetAlispTime() string {
	return r._alispTime
}

// SetAlispSign is AlispSign Setter
// 签名字符串
func (r *AlibabaalisportspassportaccountdelrelationAPIRequest) SetAlispSign(_alispSign string) error {
	r._alispSign = _alispSign
	r.Set("alisp_sign", _alispSign)
	return nil
}

// GetAlispSign AlispSign Getter
func (r AlibabaalisportspassportaccountdelrelationAPIRequest) GetAlispSign() string {
	return r._alispSign
}

// SetAppUid is AppUid Setter
// 合作方用户ID
func (r *AlibabaalisportspassportaccountdelrelationAPIRequest) SetAppUid(_appUid string) error {
	r._appUid = _appUid
	r.Set("app_uid", _appUid)
	return nil
}

// GetAppUid AppUid Getter
func (r AlibabaalisportspassportaccountdelrelationAPIRequest) GetAppUid() string {
	return r._appUid
}

// SetAliuid is Aliuid Setter
// 阿里体育会员id
func (r *AlibabaalisportspassportaccountdelrelationAPIRequest) SetAliuid(_aliuid string) error {
	r._aliuid = _aliuid
	r.Set("aliuid", _aliuid)
	return nil
}

// GetAliuid Aliuid Getter
func (r AlibabaalisportspassportaccountdelrelationAPIRequest) GetAliuid() string {
	return r._aliuid
}
