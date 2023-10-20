package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallaliautoreceiptordercheckAPIRequest 查看工单查询订单是否已付款 API请求
// tmall.aliauto.receipt.order.check
//
// 查看工单查询订单是否已付款
type TmallaliautoreceiptordercheckAPIRequest struct {
	model.Params
	// 服务商自定义门店编码
	_outerShopId string
	// 工单号
	_receiptId int64
}

// NewTmallaliautoreceiptordercheckRequest 初始化TmallaliautoreceiptordercheckAPIRequest对象
func NewTmallaliautoreceiptordercheckRequest() *TmallaliautoreceiptordercheckAPIRequest {
	return &TmallaliautoreceiptordercheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallaliautoreceiptordercheckAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.receipt.order.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallaliautoreceiptordercheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallaliautoreceiptordercheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterShopId is OuterShopId Setter
// 服务商自定义门店编码
func (r *TmallaliautoreceiptordercheckAPIRequest) SetOuterShopId(_outerShopId string) error {
	r._outerShopId = _outerShopId
	r.Set("outer_shop_id", _outerShopId)
	return nil
}

// GetOuterShopId OuterShopId Getter
func (r TmallaliautoreceiptordercheckAPIRequest) GetOuterShopId() string {
	return r._outerShopId
}

// SetReceiptId is ReceiptId Setter
// 工单号
func (r *TmallaliautoreceiptordercheckAPIRequest) SetReceiptId(_receiptId int64) error {
	r._receiptId = _receiptId
	r.Set("receipt_id", _receiptId)
	return nil
}

// GetReceiptId ReceiptId Getter
func (r TmallaliautoreceiptordercheckAPIRequest) GetReceiptId() int64 {
	return r._receiptId
}
