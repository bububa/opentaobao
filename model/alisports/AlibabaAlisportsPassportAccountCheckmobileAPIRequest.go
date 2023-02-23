package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportAccountCheckmobileAPIRequest 阿里体育会员系统--手机号验证接口 API请求
// alibaba.alisports.passport.account.checkmobile
//
// 验证三方用户的手机号，获取对应的阿里体育会员id
type AlibabaAlisportsPassportAccountCheckmobileAPIRequest struct {
	model.Params
	// 业务appkey
	_alispAppKey string
	// 调用时间戳
	_alispTime string
	// 签名字符串
	_alispSign string
	// 合作方用户ID
	_appUid string
	// 用户呢称
	_nick string
	// 手机号
	_mobile string
	// 性别 0未设置 1男 2女 3保密
	_gender string
	// 生日
	_birthday string
}

// NewAlibabaAlisportsPassportAccountCheckmobileRequest 初始化AlibabaAlisportsPassportAccountCheckmobileAPIRequest对象
func NewAlibabaAlisportsPassportAccountCheckmobileRequest() *AlibabaAlisportsPassportAccountCheckmobileAPIRequest {
	return &AlibabaAlisportsPassportAccountCheckmobileAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportAccountCheckmobileAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.passport.account.checkmobile"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportAccountCheckmobileAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlisportsPassportAccountCheckmobileAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlispAppKey is AlispAppKey Setter
// 业务appkey
func (r *AlibabaAlisportsPassportAccountCheckmobileAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// GetAlispAppKey AlispAppKey Getter
func (r AlibabaAlisportsPassportAccountCheckmobileAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// SetAlispTime is AlispTime Setter
// 调用时间戳
func (r *AlibabaAlisportsPassportAccountCheckmobileAPIRequest) SetAlispTime(_alispTime string) error {
	r._alispTime = _alispTime
	r.Set("alisp_time", _alispTime)
	return nil
}

// GetAlispTime AlispTime Getter
func (r AlibabaAlisportsPassportAccountCheckmobileAPIRequest) GetAlispTime() string {
	return r._alispTime
}

// SetAlispSign is AlispSign Setter
// 签名字符串
func (r *AlibabaAlisportsPassportAccountCheckmobileAPIRequest) SetAlispSign(_alispSign string) error {
	r._alispSign = _alispSign
	r.Set("alisp_sign", _alispSign)
	return nil
}

// GetAlispSign AlispSign Getter
func (r AlibabaAlisportsPassportAccountCheckmobileAPIRequest) GetAlispSign() string {
	return r._alispSign
}

// SetAppUid is AppUid Setter
// 合作方用户ID
func (r *AlibabaAlisportsPassportAccountCheckmobileAPIRequest) SetAppUid(_appUid string) error {
	r._appUid = _appUid
	r.Set("app_uid", _appUid)
	return nil
}

// GetAppUid AppUid Getter
func (r AlibabaAlisportsPassportAccountCheckmobileAPIRequest) GetAppUid() string {
	return r._appUid
}

// SetNick is Nick Setter
// 用户呢称
func (r *AlibabaAlisportsPassportAccountCheckmobileAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r AlibabaAlisportsPassportAccountCheckmobileAPIRequest) GetNick() string {
	return r._nick
}

// SetMobile is Mobile Setter
// 手机号
func (r *AlibabaAlisportsPassportAccountCheckmobileAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r AlibabaAlisportsPassportAccountCheckmobileAPIRequest) GetMobile() string {
	return r._mobile
}

// SetGender is Gender Setter
// 性别 0未设置 1男 2女 3保密
func (r *AlibabaAlisportsPassportAccountCheckmobileAPIRequest) SetGender(_gender string) error {
	r._gender = _gender
	r.Set("gender", _gender)
	return nil
}

// GetGender Gender Getter
func (r AlibabaAlisportsPassportAccountCheckmobileAPIRequest) GetGender() string {
	return r._gender
}

// SetBirthday is Birthday Setter
// 生日
func (r *AlibabaAlisportsPassportAccountCheckmobileAPIRequest) SetBirthday(_birthday string) error {
	r._birthday = _birthday
	r.Set("birthday", _birthday)
	return nil
}

// GetBirthday Birthday Getter
func (r AlibabaAlisportsPassportAccountCheckmobileAPIRequest) GetBirthday() string {
	return r._birthday
}
