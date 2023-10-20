package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkordersyncAPIRequest 五道口外部订单同步 API请求
// alibaba.wdk.order.sync
//
// 外部商户使用自助POS下单订单同步到五道口
type AlibabawdkordersyncAPIRequest struct {
	model.Params
	// 订单
	_receiptOrder *ReceiptOrderDo
}

// NewAlibabawdkordersyncRequest 初始化AlibabawdkordersyncAPIRequest对象
func NewAlibabawdkordersyncRequest() *AlibabawdkordersyncAPIRequest {
	return &AlibabawdkordersyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkordersyncAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.order.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkordersyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkordersyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReceiptOrder is ReceiptOrder Setter
// 订单
func (r *AlibabawdkordersyncAPIRequest) SetReceiptOrder(_receiptOrder *ReceiptOrderDo) error {
	r._receiptOrder = _receiptOrder
	r.Set("receipt_order", _receiptOrder)
	return nil
}

// GetReceiptOrder ReceiptOrder Getter
func (r AlibabawdkordersyncAPIRequest) GetReceiptOrder() *ReceiptOrderDo {
	return r._receiptOrder
}
