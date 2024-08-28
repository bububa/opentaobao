package trade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AlibabaWdkPosTradeQuery 轻pos品牌营销查询接口
// alibaba.wdk.pos.trade.query
//
// 轻pos品牌营销场景，外部商家查询营销信息
func AlibabaWdkPosTradeQuery(ctx context.Context, clt *core.SDKClient, req *trade.AlibabaWdkPosTradeQueryAPIRequest, resp *trade.AlibabaWdkPosTradeQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
