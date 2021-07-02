package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkOrderSyncAPIRequest 五道口外部订单同步 API请求
// alibaba.wdk.order.sync
//
// 外部商户使用自助POS下单订单同步到五道口
type AlibabaWdkOrderSyncAPIRequest struct {
	model.Params
	// 订单
	_receiptOrder *ReceiptOrderDo
}

// NewAlibabaWdkOrderSyncRequest 初始化AlibabaWdkOrderSyncAPIRequest对象
func NewAlibabaWdkOrderSyncRequest() *AlibabaWdkOrderSyncAPIRequest {
	return &AlibabaWdkOrderSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.order.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ReceiptOrder Setter
// 订单
func (r *AlibabaWdkOrderSyncAPIRequest) SetReceiptOrder(_receiptOrder *ReceiptOrderDo) error {
	r._receiptOrder = _receiptOrder
	r.Set("receipt_order", _receiptOrder)
	return nil
}

// Get ReceiptOrder Getter
func (r AlibabaWdkOrderSyncAPIRequest) GetReceiptOrder() *ReceiptOrderDo {
	return r._receiptOrder
}
