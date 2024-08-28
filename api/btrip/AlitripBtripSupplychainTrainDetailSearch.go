package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripSupplychainTrainDetailSearch 【商旅】火车票订单详情查询
// alitrip.btrip.supplychain.train.detail.search
//
// 【商旅】火车票订单详情查询
func AlitripBtripSupplychainTrainDetailSearch(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripSupplychainTrainDetailSearchAPIRequest, resp *btrip.AlitripBtripSupplychainTrainDetailSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
