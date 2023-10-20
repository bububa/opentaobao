package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthCodeGetcodeinfoAPIRequest 码查询功能 API请求
// alibaba.alihealth.code.getcodeinfo
//
// 码查询功能
type AlibabaAlihealthCodeGetcodeinfoAPIRequest struct {
	model.Params
	// 企业refEntId
	_refEntId string
	// 码
	_code string
}

// NewAlibabaAlihealthCodeGetcodeinfoRequest 初始化AlibabaAlihealthCodeGetcodeinfoAPIRequest对象
func NewAlibabaAlihealthCodeGetcodeinfoRequest() *AlibabaAlihealthCodeGetcodeinfoAPIRequest {
	return &AlibabaAlihealthCodeGetcodeinfoAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthCodeGetcodeinfoAPIRequest) Reset() {
	r._refEntId = ""
	r._code = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthCodeGetcodeinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.code.getcodeinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthCodeGetcodeinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthCodeGetcodeinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业refEntId
func (r *AlibabaAlihealthCodeGetcodeinfoAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthCodeGetcodeinfoAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetCode is Code Setter
// 码
func (r *AlibabaAlihealthCodeGetcodeinfoAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthCodeGetcodeinfoAPIRequest) GetCode() string {
	return r._code
}

var poolAlibabaAlihealthCodeGetcodeinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthCodeGetcodeinfoRequest()
	},
}

// GetAlibabaAlihealthCodeGetcodeinfoRequest 从 sync.Pool 获取 AlibabaAlihealthCodeGetcodeinfoAPIRequest
func GetAlibabaAlihealthCodeGetcodeinfoAPIRequest() *AlibabaAlihealthCodeGetcodeinfoAPIRequest {
	return poolAlibabaAlihealthCodeGetcodeinfoAPIRequest.Get().(*AlibabaAlihealthCodeGetcodeinfoAPIRequest)
}

// ReleaseAlibabaAlihealthCodeGetcodeinfoAPIRequest 将 AlibabaAlihealthCodeGetcodeinfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthCodeGetcodeinfoAPIRequest(v *AlibabaAlihealthCodeGetcodeinfoAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthCodeGetcodeinfoAPIRequest.Put(v)
}
