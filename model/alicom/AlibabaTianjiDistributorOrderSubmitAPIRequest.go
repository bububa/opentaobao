package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTianjiDistributorOrderSubmitAPIRequest 分销商提交受理订单 API请求
// alibaba.tianji.distributor.order.submit
//
// 分销商提交受理订单，如合约订购、充值受理等
type AlibabaTianjiDistributorOrderSubmitAPIRequest struct {
	model.Params
	// 商品编码，如手机串号
	_itemSerialNo string
	// 淘宝交易订单号
	_orderNo string
	// 供应商产品编码，如SIM卡号
	_productSerialNo string
}

// NewAlibabaTianjiDistributorOrderSubmitRequest 初始化AlibabaTianjiDistributorOrderSubmitAPIRequest对象
func NewAlibabaTianjiDistributorOrderSubmitRequest() *AlibabaTianjiDistributorOrderSubmitAPIRequest {
	return &AlibabaTianjiDistributorOrderSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTianjiDistributorOrderSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.tianji.distributor.order.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTianjiDistributorOrderSubmitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ItemSerialNo Setter
// 商品编码，如手机串号
func (r *AlibabaTianjiDistributorOrderSubmitAPIRequest) SetItemSerialNo(_itemSerialNo string) error {
	r._itemSerialNo = _itemSerialNo
	r.Set("item_serial_no", _itemSerialNo)
	return nil
}

// Get ItemSerialNo Getter
func (r AlibabaTianjiDistributorOrderSubmitAPIRequest) GetItemSerialNo() string {
	return r._itemSerialNo
}

// Set is OrderNo Setter
// 淘宝交易订单号
func (r *AlibabaTianjiDistributorOrderSubmitAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// Get OrderNo Getter
func (r AlibabaTianjiDistributorOrderSubmitAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// Set is ProductSerialNo Setter
// 供应商产品编码，如SIM卡号
func (r *AlibabaTianjiDistributorOrderSubmitAPIRequest) SetProductSerialNo(_productSerialNo string) error {
	r._productSerialNo = _productSerialNo
	r.Set("product_serial_no", _productSerialNo)
	return nil
}

// Get ProductSerialNo Getter
func (r AlibabaTianjiDistributorOrderSubmitAPIRequest) GetProductSerialNo() string {
	return r._productSerialNo
}
