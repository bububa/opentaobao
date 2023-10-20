package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatianjidistributorordersubmitAPIRequest 分销商提交受理订单 API请求
// alibaba.tianji.distributor.order.submit
//
// 分销商提交受理订单，如合约订购、充值受理等
type AlibabatianjidistributorordersubmitAPIRequest struct {
	model.Params
	// 商品编码，如手机串号
	_itemSerialNo string
	// 淘宝交易订单号
	_orderNo string
	// 供应商产品编码，如SIM卡号
	_productSerialNo string
}

// NewAlibabatianjidistributorordersubmitRequest 初始化AlibabatianjidistributorordersubmitAPIRequest对象
func NewAlibabatianjidistributorordersubmitRequest() *AlibabatianjidistributorordersubmitAPIRequest {
	return &AlibabatianjidistributorordersubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatianjidistributorordersubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.tianji.distributor.order.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatianjidistributorordersubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatianjidistributorordersubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemSerialNo is ItemSerialNo Setter
// 商品编码，如手机串号
func (r *AlibabatianjidistributorordersubmitAPIRequest) SetItemSerialNo(_itemSerialNo string) error {
	r._itemSerialNo = _itemSerialNo
	r.Set("item_serial_no", _itemSerialNo)
	return nil
}

// GetItemSerialNo ItemSerialNo Getter
func (r AlibabatianjidistributorordersubmitAPIRequest) GetItemSerialNo() string {
	return r._itemSerialNo
}

// SetOrderNo is OrderNo Setter
// 淘宝交易订单号
func (r *AlibabatianjidistributorordersubmitAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r AlibabatianjidistributorordersubmitAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// SetProductSerialNo is ProductSerialNo Setter
// 供应商产品编码，如SIM卡号
func (r *AlibabatianjidistributorordersubmitAPIRequest) SetProductSerialNo(_productSerialNo string) error {
	r._productSerialNo = _productSerialNo
	r.Set("product_serial_no", _productSerialNo)
	return nil
}

// GetProductSerialNo ProductSerialNo Getter
func (r AlibabatianjidistributorordersubmitAPIRequest) GetProductSerialNo() string {
	return r._productSerialNo
}
