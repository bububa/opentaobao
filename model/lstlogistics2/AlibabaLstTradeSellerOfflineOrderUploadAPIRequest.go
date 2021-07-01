package lstlogistics2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstTradeSellerOfflineOrderUploadAPIRequest
供应商-线下订单-导入接口 API请求
alibaba.lst.trade.seller.offline.order.upload

供应商线下订单数据上传、实现和零售通本地云仓订单的共配 */
type AlibabaLstTradeSellerOfflineOrderUploadAPIRequest struct {
	model.Params
	// 入参
	_offlineOrderUploadParam *LstOffLineOrderUploadParam
}

// New
