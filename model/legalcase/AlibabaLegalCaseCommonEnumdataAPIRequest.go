package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalCaseCommonEnumdataAPIRequest 获取通用枚举值接口 API请求
// alibaba.legal.case.common.enumdata
//
// 获取通用枚举值接口
type AlibabaLegalCaseCommonEnumdataAPIRequest struct {
	model.Params
	// bu
	_key string
	// 语言
	_lang string
}

// NewAlibabaLegalCaseCommonEnumdataRequest 初始化AlibabaLegalCaseCommonEnumdataAPIRequest对象
func NewAlibabaLegalCaseCommonEnumdataRequest() *AlibabaLegalCaseCommonEnumdataAPIRequest {
	return &AlibabaLegalCaseCommonEnumdataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseCommonEnumdataAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.case.common.enumdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseCommonEnumdataAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Key Setter
// bu
func (r *AlibabaLegalCaseCommonEnumdataAPIRequest) SetKey(_key string) error {
	r._key = _key
	r.Set("key", _key)
	return nil
}

// Get Key Getter
func (r AlibabaLegalCaseCommonEnumdataAPIRequest) GetKey() string {
	return r._key
}

// Set is Lang Setter
// 语言
func (r *AlibabaLegalCaseCommonEnumdataAPIRequest) SetLang(_lang string) error {
	r._lang = _lang
	r.Set("lang", _lang)
	return nil
}

// Get Lang Getter
func (r AlibabaLegalCaseCommonEnumdataAPIRequest) GetLang() string {
	return r._lang
}
