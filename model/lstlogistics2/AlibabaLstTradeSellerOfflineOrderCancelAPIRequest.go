package lstlogistics2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstTradeSellerOfflineOrderCancelAPIRequest
供应商-线下订单-取消接口 API请求
alibaba.lst.trade.seller.offline.order.cancel

供应商线下订单数据上传之后取消 */
type AlibabaLstTradeSellerOfflineOrderCancelAPIRequest struct {
	model.Params
	// 入参
	_offlineOrderCancalParam *LstOfflineOrderCancalParam
}

// New
