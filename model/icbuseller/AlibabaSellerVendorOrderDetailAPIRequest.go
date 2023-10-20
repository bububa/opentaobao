package icbuseller

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabasellervendororderdetailAPIRequest 国际站服务市场订单详情接口 API请求
// alibaba.seller.vendor.order.detail
//
// 国际站服务市场订单列表接口
type AlibabasellervendororderdetailAPIRequest struct {
	model.Params
	// 订单编号
	_orderNo string
}

// NewAlibabasellervendororderdetailRequest 初始化AlibabasellervendororderdetailAPIRequest对象
func NewAlibabasellervendororderdetailRequest() *AlibabasellervendororderdetailAPIRequest {
	return &AlibabasellervendororderdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabasellervendororderdetailAPIRequest) GetApiMethodName() string {
	return "alibaba.seller.vendor.order.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabasellervendororderdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabasellervendororderdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderNo is OrderNo Setter
// 订单编号
func (r *AlibabasellervendororderdetailAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r AlibabasellervendororderdetailAPIRequest) GetOrderNo() string {
	return r._orderNo
}
