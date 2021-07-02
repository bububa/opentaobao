package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthNrRxQueryimageAPIRequest o2o查看处方图片 API请求
// alibaba.alihealth.nr.rx.queryimage
//
// o2o商家查看处方图片，包括电子图片与纸质图片
type AlibabaAlihealthNrRxQueryimageAPIRequest struct {
	model.Params
	// 订单编号
	_orderId int64
}

// NewAlibabaAlihealthNrRxQueryimageRequest 初始化AlibabaAlihealthNrRxQueryimageAPIRequest对象
func NewAlibabaAlihealthNrRxQueryimageRequest() *AlibabaAlihealthNrRxQueryimageAPIRequest {
	return &AlibabaAlihealthNrRxQueryimageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthNrRxQueryimageAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.nr.rx.queryimage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthNrRxQueryimageAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 订单编号
func (r *AlibabaAlihealthNrRxQueryimageAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r AlibabaAlihealthNrRxQueryimageAPIRequest) GetOrderId() int64 {
	return r._orderId
}
