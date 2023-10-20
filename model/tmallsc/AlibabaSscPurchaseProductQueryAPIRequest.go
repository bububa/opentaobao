package tmallsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscPurchaseProductQueryAPIRequest 查询已采购的服务产品 API请求
// alibaba.ssc.purchase.product.query
//
// 查询已采购的服务产品
type AlibabaSscPurchaseProductQueryAPIRequest struct {
	model.Params
}

// NewAlibabaSscPurchaseProductQueryRequest 初始化AlibabaSscPurchaseProductQueryAPIRequest对象
func NewAlibabaSscPurchaseProductQueryRequest() *AlibabaSscPurchaseProductQueryAPIRequest {
	return &AlibabaSscPurchaseProductQueryAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSscPurchaseProductQueryAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscPurchaseProductQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.purchase.product.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscPurchaseProductQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSscPurchaseProductQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaSscPurchaseProductQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSscPurchaseProductQueryRequest()
	},
}

// GetAlibabaSscPurchaseProductQueryRequest 从 sync.Pool 获取 AlibabaSscPurchaseProductQueryAPIRequest
func GetAlibabaSscPurchaseProductQueryAPIRequest() *AlibabaSscPurchaseProductQueryAPIRequest {
	return poolAlibabaSscPurchaseProductQueryAPIRequest.Get().(*AlibabaSscPurchaseProductQueryAPIRequest)
}

// ReleaseAlibabaSscPurchaseProductQueryAPIRequest 将 AlibabaSscPurchaseProductQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaSscPurchaseProductQueryAPIRequest(v *AlibabaSscPurchaseProductQueryAPIRequest) {
	v.Reset()
	poolAlibabaSscPurchaseProductQueryAPIRequest.Put(v)
}
