package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaPricePromotionActivityQuery 查询盒马帮档期活动详情
// alibaba.price.promotion.activity.query
//
// 查询盒马帮档期活动详情
func AlibabaPricePromotionActivityQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaPricePromotionActivityQueryAPIRequest, resp *wdk.AlibabaPricePromotionActivityQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
