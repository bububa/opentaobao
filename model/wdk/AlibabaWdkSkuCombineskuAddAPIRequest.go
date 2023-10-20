package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuCombineskuAddAPIRequest 组合商品新增接口 API请求
// alibaba.wdk.sku.combinesku.add
//
// 组合商品新增接口
type AlibabaWdkSkuCombineskuAddAPIRequest struct {
	model.Params
	// 请求参数
	_paramList []SkuDo
}

// NewAlibabaWdkSkuCombineskuAddRequest 初始化AlibabaWdkSkuCombineskuAddAPIRequest对象
func NewAlibabaWdkSkuCombineskuAddRequest() *AlibabaWdkSkuCombineskuAddAPIRequest {
	return &AlibabaWdkSkuCombineskuAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSkuCombineskuAddAPIRequest) Reset() {
	r._paramList = r._paramList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuCombineskuAddAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.combinesku.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuCombineskuAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSkuCombineskuAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 请求参数
func (r *AlibabaWdkSkuCombineskuAddAPIRequest) SetParamList(_paramList []SkuDo) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r AlibabaWdkSkuCombineskuAddAPIRequest) GetParamList() []SkuDo {
	return r._paramList
}

var poolAlibabaWdkSkuCombineskuAddAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSkuCombineskuAddRequest()
	},
}

// GetAlibabaWdkSkuCombineskuAddRequest 从 sync.Pool 获取 AlibabaWdkSkuCombineskuAddAPIRequest
func GetAlibabaWdkSkuCombineskuAddAPIRequest() *AlibabaWdkSkuCombineskuAddAPIRequest {
	return poolAlibabaWdkSkuCombineskuAddAPIRequest.Get().(*AlibabaWdkSkuCombineskuAddAPIRequest)
}

// ReleaseAlibabaWdkSkuCombineskuAddAPIRequest 将 AlibabaWdkSkuCombineskuAddAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSkuCombineskuAddAPIRequest(v *AlibabaWdkSkuCombineskuAddAPIRequest) {
	v.Reset()
	poolAlibabaWdkSkuCombineskuAddAPIRequest.Put(v)
}
