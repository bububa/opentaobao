package tmallsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscPurchaseServicedefinitionParamQueryAPIRequest 查询采购服务定义参数信息 API请求
// alibaba.ssc.purchase.servicedefinition.param.query
//
// 查询采购服务定义参数信息
type AlibabaSscPurchaseServicedefinitionParamQueryAPIRequest struct {
	model.Params
	// 服务产品id
	_productId int64
}

// NewAlibabaSscPurchaseServicedefinitionParamQueryRequest 初始化AlibabaSscPurchaseServicedefinitionParamQueryAPIRequest对象
func NewAlibabaSscPurchaseServicedefinitionParamQueryRequest() *AlibabaSscPurchaseServicedefinitionParamQueryAPIRequest {
	return &AlibabaSscPurchaseServicedefinitionParamQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSscPurchaseServicedefinitionParamQueryAPIRequest) Reset() {
	r._productId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSscPurchaseServicedefinitionParamQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ssc.purchase.servicedefinition.param.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSscPurchaseServicedefinitionParamQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSscPurchaseServicedefinitionParamQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 服务产品id
func (r *AlibabaSscPurchaseServicedefinitionParamQueryAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r AlibabaSscPurchaseServicedefinitionParamQueryAPIRequest) GetProductId() int64 {
	return r._productId
}

var poolAlibabaSscPurchaseServicedefinitionParamQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSscPurchaseServicedefinitionParamQueryRequest()
	},
}

// GetAlibabaSscPurchaseServicedefinitionParamQueryRequest 从 sync.Pool 获取 AlibabaSscPurchaseServicedefinitionParamQueryAPIRequest
func GetAlibabaSscPurchaseServicedefinitionParamQueryAPIRequest() *AlibabaSscPurchaseServicedefinitionParamQueryAPIRequest {
	return poolAlibabaSscPurchaseServicedefinitionParamQueryAPIRequest.Get().(*AlibabaSscPurchaseServicedefinitionParamQueryAPIRequest)
}

// ReleaseAlibabaSscPurchaseServicedefinitionParamQueryAPIRequest 将 AlibabaSscPurchaseServicedefinitionParamQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaSscPurchaseServicedefinitionParamQueryAPIRequest(v *AlibabaSscPurchaseServicedefinitionParamQueryAPIRequest) {
	v.Reset()
	poolAlibabaSscPurchaseServicedefinitionParamQueryAPIRequest.Put(v)
}
