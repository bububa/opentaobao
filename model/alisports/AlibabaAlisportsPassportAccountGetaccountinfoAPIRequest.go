package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalisportspassportaccountgetaccountinfoAPIRequest 获取会员信息 API请求
// alibaba.alisports.passport.account.getaccountinfo
//
// 获取阿里体育会员信息
type AlibabaalisportspassportaccountgetaccountinfoAPIRequest struct {
	model.Params
	// 当前时间戳
	_alispTime string
	// 业务appkey
	_alispAppKey string
	// 业务加密参数
	_alispSign string
	// 要查询的值
	_value string
	// 决定返回值是否包含扩展字段
	_extInfoType string
	// 是否获取详情0否1是 默认0
	_needDetail int64
	// 查询类型：1.用户的阿里体育id, 4.用户通过登录生成的sso_token
	_type int64
}

// NewAlibabaalisportspassportaccountgetaccountinfoRequest 初始化AlibabaalisportspassportaccountgetaccountinfoAPIRequest对象
func NewAlibabaalisportspassportaccountgetaccountinfoRequest() *AlibabaalisportspassportaccountgetaccountinfoAPIRequest {
	return &AlibabaalisportspassportaccountgetaccountinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalisportspassportaccountgetaccountinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.passport.account.getaccountinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalisportspassportaccountgetaccountinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalisportspassportaccountgetaccountinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAlispTime is AlispTime Setter
// 当前时间戳
func (r *AlibabaalisportspassportaccountgetaccountinfoAPIRequest) SetAlispTime(_alispTime string) error {
	r._alispTime = _alispTime
	r.Set("alisp_time", _alispTime)
	return nil
}

// GetAlispTime AlispTime Getter
func (r AlibabaalisportspassportaccountgetaccountinfoAPIRequest) GetAlispTime() string {
	return r._alispTime
}

// SetAlispAppKey is AlispAppKey Setter
// 业务appkey
func (r *AlibabaalisportspassportaccountgetaccountinfoAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// GetAlispAppKey AlispAppKey Getter
func (r AlibabaalisportspassportaccountgetaccountinfoAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// SetAlispSign is AlispSign Setter
// 业务加密参数
func (r *AlibabaalisportspassportaccountgetaccountinfoAPIRequest) SetAlispSign(_alispSign string) error {
	r._alispSign = _alispSign
	r.Set("alisp_sign", _alispSign)
	return nil
}

// GetAlispSign AlispSign Getter
func (r AlibabaalisportspassportaccountgetaccountinfoAPIRequest) GetAlispSign() string {
	return r._alispSign
}

// SetValue is Value Setter
// 要查询的值
func (r *AlibabaalisportspassportaccountgetaccountinfoAPIRequest) SetValue(_value string) error {
	r._value = _value
	r.Set("value", _value)
	return nil
}

// GetValue Value Getter
func (r AlibabaalisportspassportaccountgetaccountinfoAPIRequest) GetValue() string {
	return r._value
}

// SetExtInfoType is ExtInfoType Setter
// 决定返回值是否包含扩展字段
func (r *AlibabaalisportspassportaccountgetaccountinfoAPIRequest) SetExtInfoType(_extInfoType string) error {
	r._extInfoType = _extInfoType
	r.Set("ext_info_type", _extInfoType)
	return nil
}

// GetExtInfoType ExtInfoType Getter
func (r AlibabaalisportspassportaccountgetaccountinfoAPIRequest) GetExtInfoType() string {
	return r._extInfoType
}

// SetNeedDetail is NeedDetail Setter
// 是否获取详情0否1是 默认0
func (r *AlibabaalisportspassportaccountgetaccountinfoAPIRequest) SetNeedDetail(_needDetail int64) error {
	r._needDetail = _needDetail
	r.Set("need_detail", _needDetail)
	return nil
}

// GetNeedDetail NeedDetail Getter
func (r AlibabaalisportspassportaccountgetaccountinfoAPIRequest) GetNeedDetail() int64 {
	return r._needDetail
}

// SetType is Type Setter
// 查询类型：1.用户的阿里体育id, 4.用户通过登录生成的sso_token
func (r *AlibabaalisportspassportaccountgetaccountinfoAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaalisportspassportaccountgetaccountinfoAPIRequest) GetType() int64 {
	return r._type
}
