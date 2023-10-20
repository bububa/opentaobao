package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuCombineskuUpdateAPIRequest 组合商品更新接口 API请求
// alibaba.wdk.sku.combinesku.update
//
// 组合商品修改接口
type AlibabaWdkSkuCombineskuUpdateAPIRequest struct {
	model.Params
	// 请求参数
	_paramList []SkuDo
}

// NewAlibabaWdkSkuCombineskuUpdateRequest 初始化AlibabaWdkSkuCombineskuUpdateAPIRequest对象
func NewAlibabaWdkSkuCombineskuUpdateRequest() *AlibabaWdkSkuCombineskuUpdateAPIRequest {
	return &AlibabaWdkSkuCombineskuUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSkuCombineskuUpdateAPIRequest) Reset() {
	r._paramList = r._paramList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuCombineskuUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.combinesku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuCombineskuUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSkuCombineskuUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 请求参数
func (r *AlibabaWdkSkuCombineskuUpdateAPIRequest) SetParamList(_paramList []SkuDo) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r AlibabaWdkSkuCombineskuUpdateAPIRequest) GetParamList() []SkuDo {
	return r._paramList
}

var poolAlibabaWdkSkuCombineskuUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSkuCombineskuUpdateRequest()
	},
}

// GetAlibabaWdkSkuCombineskuUpdateRequest 从 sync.Pool 获取 AlibabaWdkSkuCombineskuUpdateAPIRequest
func GetAlibabaWdkSkuCombineskuUpdateAPIRequest() *AlibabaWdkSkuCombineskuUpdateAPIRequest {
	return poolAlibabaWdkSkuCombineskuUpdateAPIRequest.Get().(*AlibabaWdkSkuCombineskuUpdateAPIRequest)
}

// ReleaseAlibabaWdkSkuCombineskuUpdateAPIRequest 将 AlibabaWdkSkuCombineskuUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSkuCombineskuUpdateAPIRequest(v *AlibabaWdkSkuCombineskuUpdateAPIRequest) {
	v.Reset()
	poolAlibabaWdkSkuCombineskuUpdateAPIRequest.Put(v)
}
