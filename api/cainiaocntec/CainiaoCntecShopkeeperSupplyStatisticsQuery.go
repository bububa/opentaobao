package cainiaocntec

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaocntec"
)

// CainiaoCntecShopkeeperSupplyStatisticsQuery 团购业务供货商查询门店统计数据
// cainiao.cntec.shopkeeper.supply.statistics.query
//
// 查询门店售卖商品统计数据
func CainiaoCntecShopkeeperSupplyStatisticsQuery(ctx context.Context, clt *core.SDKClient, req *cainiaocntec.CainiaoCntecShopkeeperSupplyStatisticsQueryAPIRequest, resp *cainiaocntec.CainiaoCntecShopkeeperSupplyStatisticsQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
