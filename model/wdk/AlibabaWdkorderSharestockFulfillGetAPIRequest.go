package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkorderSharestockFulfillGetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdkorder.sharestock.fulfill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkorderSharestockFulfillGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
