package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkOrderSyncAPIRequest
五道口外部订单同步 API请求
alibaba.wdk.order.sync

外部商户使用自助POS下单订单同步到五道口 */
type AlibabaWdkOrderSyncAPIRequest struct {
	model.Params
	// 订单
	_receiptOrder *ReceiptOrderDo
}

// New
