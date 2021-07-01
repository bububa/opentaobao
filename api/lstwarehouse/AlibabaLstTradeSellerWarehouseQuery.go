package lstwarehouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstwarehouse"
)

/* AlibabaLstTradeSellerWarehouseQuery
供应商-本云商家-仓库查询接口
alibaba.lst.trade.seller.warehouse.query

查询本地云仓商家的仓库 */
func AlibabaLstTradeSellerWarehouseQuery(clt *core.SDKClient, req *lstwarehouse.AlibabaLstTradeSellerWarehouseQueryAPIRequest, session string) (*lstwarehouse.AlibabaLstTradeSellerWarehouseQueryAPIResponse, error) {
	var resp lstwarehouse.AlibabaLstTradeSellerWarehouseQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
