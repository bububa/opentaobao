package lstwarehouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstwarehouse"
)

// AlibabaLstTradeSellerWarehouseQuery 供应商-本云商家-仓库查询接口
// alibaba.lst.trade.seller.warehouse.query
//
// 查询本地云仓商家的仓库
func AlibabaLstTradeSellerWarehouseQuery(ctx context.Context, clt *core.SDKClient, req *lstwarehouse.AlibabaLstTradeSellerWarehouseQueryAPIRequest, resp *lstwarehouse.AlibabaLstTradeSellerWarehouseQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
