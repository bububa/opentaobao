package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuQueryAPIRequest 查询商品 API请求
// alibaba.wdk.sku.query
//
// 查询商品
type AlibabaWdkSkuQueryAPIRequest struct {
	model.Params
	// 入参
	_param *SkuQueryDo
}

// NewAlibabaWdkSkuQueryRequest 初始化AlibabaWdkSkuQueryAPIRequest对象
func NewAlibabaWdkSkuQueryRequest() *AlibabaWdkSkuQueryAPIRequest {
	return &AlibabaWdkSkuQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkSkuQueryAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkSkuQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.sku.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkSkuQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkSkuQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabaWdkSkuQueryAPIRequest) SetParam(_param *SkuQueryDo) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkSkuQueryAPIRequest) GetParam() *SkuQueryDo {
	return r._param
}

var poolAlibabaWdkSkuQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkSkuQueryRequest()
	},
}

// GetAlibabaWdkSkuQueryRequest 从 sync.Pool 获取 AlibabaWdkSkuQueryAPIRequest
func GetAlibabaWdkSkuQueryAPIRequest() *AlibabaWdkSkuQueryAPIRequest {
	return poolAlibabaWdkSkuQueryAPIRequest.Get().(*AlibabaWdkSkuQueryAPIRequest)
}

// ReleaseAlibabaWdkSkuQueryAPIRequest 将 AlibabaWdkSkuQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkSkuQueryAPIRequest(v *AlibabaWdkSkuQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkSkuQueryAPIRequest.Put(v)
}
