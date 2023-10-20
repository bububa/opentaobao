package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalcasecommonenumdataAPIRequest 获取通用枚举值接口 API请求
// alibaba.legal.case.common.enumdata
//
// 获取通用枚举值接口
type AlibabalegalcasecommonenumdataAPIRequest struct {
	model.Params
	// bu
	_key string
	// 语言
	_lang string
}

// NewAlibabalegalcasecommonenumdataRequest 初始化AlibabalegalcasecommonenumdataAPIRequest对象
func NewAlibabalegalcasecommonenumdataRequest() *AlibabalegalcasecommonenumdataAPIRequest {
	return &AlibabalegalcasecommonenumdataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalcasecommonenumdataAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.case.common.enumdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalcasecommonenumdataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalcasecommonenumdataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKey is Key Setter
// bu
func (r *AlibabalegalcasecommonenumdataAPIRequest) SetKey(_key string) error {
	r._key = _key
	r.Set("key", _key)
	return nil
}

// GetKey Key Getter
func (r AlibabalegalcasecommonenumdataAPIRequest) GetKey() string {
	return r._key
}

// SetLang is Lang Setter
// 语言
func (r *AlibabalegalcasecommonenumdataAPIRequest) SetLang(_lang string) error {
	r._lang = _lang
	r.Set("lang", _lang)
	return nil
}

// GetLang Lang Getter
func (r AlibabalegalcasecommonenumdataAPIRequest) GetLang() string {
	return r._lang
}
