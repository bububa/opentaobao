package paimai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// TaobaoPaimaiNftOrderinfoQuery 查询订单类型
// taobao.paimai.nft.orderinfo.query
//
// 查询订单类型
func TaobaoPaimaiNftOrderinfoQuery(ctx context.Context, clt *core.SDKClient, req *paimai.TaobaoPaimaiNftOrderinfoQueryAPIRequest, resp *paimai.TaobaoPaimaiNftOrderinfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
