package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpPurchasePriceCreateAPIRequest ascp采购价写入接口 API请求
// alibaba.ascp.purchase.price.create
//
// 供应链平台采购价创建或修改接口
type AlibabaAscpPurchasePriceCreateAPIRequest struct {
	model.Params
	// 采购价创建/更新请求
	_createRequest *AlibabaAscpPurchasePriceCreateRequest
}

// NewAlibabaAscpPurchasePriceCreateRequest 初始化AlibabaAscpPurchasePriceCreateAPIRequest对象
func NewAlibabaAscpPurchasePriceCreateRequest() *AlibabaAscpPurchasePriceCreateAPIRequest {
	return &AlibabaAscpPurchasePriceCreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpPurchasePriceCreateAPIRequest) Reset() {
	r._createRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpPurchasePriceCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.purchase.price.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpPurchasePriceCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpPurchasePriceCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreateRequest is CreateRequest Setter
// 采购价创建/更新请求
func (r *AlibabaAscpPurchasePriceCreateAPIRequest) SetCreateRequest(_createRequest *AlibabaAscpPurchasePriceCreateRequest) error {
	r._createRequest = _createRequest
	r.Set("create_request", _createRequest)
	return nil
}

// GetCreateRequest CreateRequest Getter
func (r AlibabaAscpPurchasePriceCreateAPIRequest) GetCreateRequest() *AlibabaAscpPurchasePriceCreateRequest {
	return r._createRequest
}

var poolAlibabaAscpPurchasePriceCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpPurchasePriceCreateRequest()
	},
}

// GetAlibabaAscpPurchasePriceCreateRequest 从 sync.Pool 获取 AlibabaAscpPurchasePriceCreateAPIRequest
func GetAlibabaAscpPurchasePriceCreateAPIRequest() *AlibabaAscpPurchasePriceCreateAPIRequest {
	return poolAlibabaAscpPurchasePriceCreateAPIRequest.Get().(*AlibabaAscpPurchasePriceCreateAPIRequest)
}

// ReleaseAlibabaAscpPurchasePriceCreateAPIRequest 将 AlibabaAscpPurchasePriceCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpPurchasePriceCreateAPIRequest(v *AlibabaAscpPurchasePriceCreateAPIRequest) {
	v.Reset()
	poolAlibabaAscpPurchasePriceCreateAPIRequest.Put(v)
}
