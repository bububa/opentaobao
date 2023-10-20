package lstlogistics2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics2"
)

// AlibabaLstTradeSellerOfflineOrderCancel 供应商-线下订单-取消接口
// alibaba.lst.trade.seller.offline.order.cancel
//
// 供应商线下订单数据上传之后取消
func AlibabaLstTradeSellerOfflineOrderCancel(clt *core.SDKClient, req *lstlogistics2.AlibabaLstTradeSellerOfflineOrderCancelAPIRequest, session string) (*lstlogistics2.AlibabaLstTradeSellerOfflineOrderCancelAPIResponse, error) {
	var resp lstlogistics2.AlibabaLstTradeSellerOfflineOrderCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
