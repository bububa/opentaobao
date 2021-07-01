package lstlogistics2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstTradeSellerOfflineOrderQueryAPIRequest
供应商-线下订单-查询接口 API请求
alibaba.lst.trade.seller.offline.order.query

供应商线下订单数据上传后查询物流状态 */
type AlibabaLstTradeSellerOfflineOrderQueryAPIRequest struct {
	model.Params
	// 入参
	_offlineOrderQueryParam *LstOfflineOrderQueryParam
}

// New
