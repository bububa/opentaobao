package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkOrderSyncAPIRequest) Reset() {
	r._receiptOrder = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.order.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkOrderSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReceiptOrder is ReceiptOrder Setter
// 订单
func (r *AlibabaWdkOrderSyncAPIRequest) SetReceiptOrder(_receiptOrder *ReceiptOrderDo) error {
	r._receiptOrder = _receiptOrder
	r.Set("receipt_order", _receiptOrder)
	return nil
}

// GetReceiptOrder ReceiptOrder Getter
func (r AlibabaWdkOrderSyncAPIRequest) GetReceiptOrder() *ReceiptOrderDo {
	return r._receiptOrder
}

var poolAlibabaWdkOrderSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkOrderSyncRequest()
	},
}

// GetAlibabaWdkOrderSyncRequest 从 sync.Pool 获取 AlibabaWdkOrderSyncAPIRequest
func GetAlibabaWdkOrderSyncAPIRequest() *AlibabaWdkOrderSyncAPIRequest {
	return poolAlibabaWdkOrderSyncAPIRequest.Get().(*AlibabaWdkOrderSyncAPIRequest)
}

// ReleaseAlibabaWdkOrderSyncAPIRequest 将 AlibabaWdkOrderSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkOrderSyncAPIRequest(v *AlibabaWdkOrderSyncAPIRequest) {
	v.Reset()
	poolAlibabaWdkOrderSyncAPIRequest.Put(v)
}
