package ascpchannel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopCnReverseWarehouseorderUpdateAPIRequest 供应链中台逆向入库单修改服务 API请求
// alibaba.ascp.uop.cn.reverse.warehouseorder.update
//
// 供应链中台逆向入库单修改服务
type AlibabaAscpUopCnReverseWarehouseorderUpdateAPIRequest struct {
	model.Params
	// 逆向入库单号
	_orderCode string
	// 退款原因
	_refundReason string
	// 是否已经退款
	_refunded bool
}

// NewAlibabaAscpUopCnReverseWarehouseorderUpdateRequest 初始化AlibabaAscpUopCnReverseWarehouseorderUpdateAPIRequest对象
func NewAlibabaAscpUopCnReverseWarehouseorderUpdateRequest() *AlibabaAscpUopCnReverseWarehouseorderUpdateAPIRequest {
	return &AlibabaAscpUopCnReverseWarehouseorderUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopCnReverseWarehouseorderUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.cn.reverse.warehouseorder.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopCnReverseWarehouseorderUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpUopCnReverseWarehouseorderUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// 逆向入库单号
func (r *AlibabaAscpUopCnReverseWarehouseorderUpdateAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r AlibabaAscpUopCnReverseWarehouseorderUpdateAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetRefundReason is RefundReason Setter
// 退款原因
func (r *AlibabaAscpUopCnReverseWarehouseorderUpdateAPIRequest) SetRefundReason(_refundReason string) error {
	r._refundReason = _refundReason
	r.Set("refund_reason", _refundReason)
	return nil
}

// GetRefundReason RefundReason Getter
func (r AlibabaAscpUopCnReverseWarehouseorderUpdateAPIRequest) GetRefundReason() string {
	return r._refundReason
}

// SetRefunded is Refunded Setter
// 是否已经退款
func (r *AlibabaAscpUopCnReverseWarehouseorderUpdateAPIRequest) SetRefunded(_refunded bool) error {
	r._refunded = _refunded
	r.Set("refunded", _refunded)
	return nil
}

// GetRefunded Refunded Getter
func (r AlibabaAscpUopCnReverseWarehouseorderUpdateAPIRequest) GetRefunded() bool {
	return r._refunded
}
