package alicom

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTianjiDistributorOrderSubmitAPIRequest) Reset() {
	r._itemSerialNo = ""
	r._orderNo = ""
	r._productSerialNo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTianjiDistributorOrderSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.tianji.distributor.order.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTianjiDistributorOrderSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTianjiDistributorOrderSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemSerialNo is ItemSerialNo Setter
// 商品编码，如手机串号
func (r *AlibabaTianjiDistributorOrderSubmitAPIRequest) SetItemSerialNo(_itemSerialNo string) error {
	r._itemSerialNo = _itemSerialNo
	r.Set("item_serial_no", _itemSerialNo)
	return nil
}

// GetItemSerialNo ItemSerialNo Getter
func (r AlibabaTianjiDistributorOrderSubmitAPIRequest) GetItemSerialNo() string {
	return r._itemSerialNo
}

// SetOrderNo is OrderNo Setter
// 淘宝交易订单号
func (r *AlibabaTianjiDistributorOrderSubmitAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r AlibabaTianjiDistributorOrderSubmitAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// SetProductSerialNo is ProductSerialNo Setter
// 供应商产品编码，如SIM卡号
func (r *AlibabaTianjiDistributorOrderSubmitAPIRequest) SetProductSerialNo(_productSerialNo string) error {
	r._productSerialNo = _productSerialNo
	r.Set("product_serial_no", _productSerialNo)
	return nil
}

// GetProductSerialNo ProductSerialNo Getter
func (r AlibabaTianjiDistributorOrderSubmitAPIRequest) GetProductSerialNo() string {
	return r._productSerialNo
}

var poolAlibabaTianjiDistributorOrderSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTianjiDistributorOrderSubmitRequest()
	},
}

// GetAlibabaTianjiDistributorOrderSubmitRequest 从 sync.Pool 获取 AlibabaTianjiDistributorOrderSubmitAPIRequest
func GetAlibabaTianjiDistributorOrderSubmitAPIRequest() *AlibabaTianjiDistributorOrderSubmitAPIRequest {
	return poolAlibabaTianjiDistributorOrderSubmitAPIRequest.Get().(*AlibabaTianjiDistributorOrderSubmitAPIRequest)
}

// ReleaseAlibabaTianjiDistributorOrderSubmitAPIRequest 将 AlibabaTianjiDistributorOrderSubmitAPIRequest 放入 sync.Pool
func ReleaseAlibabaTianjiDistributorOrderSubmitAPIRequest(v *AlibabaTianjiDistributorOrderSubmitAPIRequest) {
	v.Reset()
	poolAlibabaTianjiDistributorOrderSubmitAPIRequest.Put(v)
}
