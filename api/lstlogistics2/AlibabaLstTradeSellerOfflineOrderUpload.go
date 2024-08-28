package lstlogistics2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics2"
)

// AlibabaLstTradeSellerOfflineOrderUpload 供应商-线下订单-导入接口
// alibaba.lst.trade.seller.offline.order.upload
//
// 供应商线下订单数据上传、实现和零售通本地云仓订单的共配
func AlibabaLstTradeSellerOfflineOrderUpload(ctx context.Context, clt *core.SDKClient, req *lstlogistics2.AlibabaLstTradeSellerOfflineOrderUploadAPIRequest, resp *lstlogistics2.AlibabaLstTradeSellerOfflineOrderUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
