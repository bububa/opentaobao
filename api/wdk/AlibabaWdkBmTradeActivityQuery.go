package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkBmTradeActivityQuery 品牌营销的订单活动信息查询
// alibaba.wdk.bm.trade.activity.query
//
// 品牌营销的订单活动信息查询
func AlibabaWdkBmTradeActivityQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkBmTradeActivityQueryAPIRequest, resp *wdk.AlibabaWdkBmTradeActivityQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
