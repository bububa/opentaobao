package legalcase

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalCaseCommonEnumdataAPIRequest) Reset() {
	r._key = ""
	r._lang = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseCommonEnumdataAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.case.common.enumdata"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseCommonEnumdataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalCaseCommonEnumdataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKey is Key Setter
// bu
func (r *AlibabaLegalCaseCommonEnumdataAPIRequest) SetKey(_key string) error {
	r._key = _key
	r.Set("key", _key)
	return nil
}

// GetKey Key Getter
func (r AlibabaLegalCaseCommonEnumdataAPIRequest) GetKey() string {
	return r._key
}

// SetLang is Lang Setter
// 语言
func (r *AlibabaLegalCaseCommonEnumdataAPIRequest) SetLang(_lang string) error {
	r._lang = _lang
	r.Set("lang", _lang)
	return nil
}

// GetLang Lang Getter
func (r AlibabaLegalCaseCommonEnumdataAPIRequest) GetLang() string {
	return r._lang
}

var poolAlibabaLegalCaseCommonEnumdataAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalCaseCommonEnumdataRequest()
	},
}

// GetAlibabaLegalCaseCommonEnumdataRequest 从 sync.Pool 获取 AlibabaLegalCaseCommonEnumdataAPIRequest
func GetAlibabaLegalCaseCommonEnumdataAPIRequest() *AlibabaLegalCaseCommonEnumdataAPIRequest {
	return poolAlibabaLegalCaseCommonEnumdataAPIRequest.Get().(*AlibabaLegalCaseCommonEnumdataAPIRequest)
}

// ReleaseAlibabaLegalCaseCommonEnumdataAPIRequest 将 AlibabaLegalCaseCommonEnumdataAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalCaseCommonEnumdataAPIRequest(v *AlibabaLegalCaseCommonEnumdataAPIRequest) {
	v.Reset()
	poolAlibabaLegalCaseCommonEnumdataAPIRequest.Put(v)
}
