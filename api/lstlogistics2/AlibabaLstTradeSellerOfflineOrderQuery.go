package lstlogistics2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics2"
)

// AlibabaLstTradeSellerOfflineOrderQuery 供应商-线下订单-查询接口
// alibaba.lst.trade.seller.offline.order.query
//
// 供应商线下订单数据上传后查询物流状态
func AlibabaLstTradeSellerOfflineOrderQuery(ctx context.Context, clt *core.SDKClient, req *lstlogistics2.AlibabaLstTradeSellerOfflineOrderQueryAPIRequest, resp *lstlogistics2.AlibabaLstTradeSellerOfflineOrderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
