package lstlogistics2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics2"
)

// AlibabaLstTradeSellerOfflineOrderQuery 供应商-线下订单-查询接口
// alibaba.lst.trade.seller.offline.order.query
//
// 供应商线下订单数据上传后查询物流状态
func AlibabaLstTradeSellerOfflineOrderQuery(clt *core.SDKClient, req *lstlogistics2.AlibabaLstTradeSellerOfflineOrderQueryAPIRequest, session string) (*lstlogistics2.AlibabaLstTradeSellerOfflineOrderQueryAPIResponse, error) {
	var resp lstlogistics2.AlibabaLstTradeSellerOfflineOrderQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
