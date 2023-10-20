package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkordersharestockfulfillgetAPIRequest 商户订单履约数据获取 API请求
// alibaba.wdkorder.sharestock.fulfill.get
//
// 商户订单履约数据获取
type AlibabawdkordersharestockfulfillgetAPIRequest struct {
	model.Params
	// 履约单ID
	_fulfillOrderId string
}

// NewAlibabawdkordersharestockfulfillgetRequest 初始化AlibabawdkordersharestockfulfillgetAPIRequest对象
func NewAlibabawdkordersharestockfulfillgetRequest() *AlibabawdkordersharestockfulfillgetAPIRequest {
	return &AlibabawdkordersharestockfulfillgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkordersharestockfulfillgetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdkorder.sharestock.fulfill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkordersharestockfulfillgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkordersharestockfulfillgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFulfillOrderId is FulfillOrderId Setter
// 履约单ID
func (r *AlibabawdkordersharestockfulfillgetAPIRequest) SetFulfillOrderId(_fulfillOrderId string) error {
	r._fulfillOrderId = _fulfillOrderId
	r.Set("fulfill_order_id", _fulfillOrderId)
	return nil
}

// GetFulfillOrderId FulfillOrderId Getter
func (r AlibabawdkordersharestockfulfillgetAPIRequest) GetFulfillOrderId() string {
	return r._fulfillOrderId
}
