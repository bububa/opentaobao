package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkorderSharestockFulfillGetAPIRequest 商户订单履约数据获取 API请求
// alibaba.wdkorder.sharestock.fulfill.get
//
// 商户订单履约数据获取
type AlibabaWdkorderSharestockFulfillGetAPIRequest struct {
	model.Params
	// 履约单ID
	_fulfillOrderId string
}

// NewAlibabaWdkorderSharestockFulfillGetRequest 初始化AlibabaWdkorderSharestockFulfillGetAPIRequest对象
func NewAlibabaWdkorderSharestockFulfillGetRequest() *AlibabaWdkorderSharestockFulfillGetAPIRequest {
	return &AlibabaWdkorderSharestockFulfillGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkorderSharestockFulfillGetAPIRequest) Reset() {
	r._fulfillOrderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkorderSharestockFulfillGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdkorder.sharestock.fulfill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkorderSharestockFulfillGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkorderSharestockFulfillGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFulfillOrderId is FulfillOrderId Setter
// 履约单ID
func (r *AlibabaWdkorderSharestockFulfillGetAPIRequest) SetFulfillOrderId(_fulfillOrderId string) error {
	r._fulfillOrderId = _fulfillOrderId
	r.Set("fulfill_order_id", _fulfillOrderId)
	return nil
}

// GetFulfillOrderId FulfillOrderId Getter
func (r AlibabaWdkorderSharestockFulfillGetAPIRequest) GetFulfillOrderId() string {
	return r._fulfillOrderId
}

var poolAlibabaWdkorderSharestockFulfillGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkorderSharestockFulfillGetRequest()
	},
}

// GetAlibabaWdkorderSharestockFulfillGetRequest 从 sync.Pool 获取 AlibabaWdkorderSharestockFulfillGetAPIRequest
func GetAlibabaWdkorderSharestockFulfillGetAPIRequest() *AlibabaWdkorderSharestockFulfillGetAPIRequest {
	return poolAlibabaWdkorderSharestockFulfillGetAPIRequest.Get().(*AlibabaWdkorderSharestockFulfillGetAPIRequest)
}

// ReleaseAlibabaWdkorderSharestockFulfillGetAPIRequest 将 AlibabaWdkorderSharestockFulfillGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkorderSharestockFulfillGetAPIRequest(v *AlibabaWdkorderSharestockFulfillGetAPIRequest) {
	v.Reset()
	poolAlibabaWdkorderSharestockFulfillGetAPIRequest.Put(v)
}
