package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// AlibabaLatourStrategyShow 阿里巴巴权益投放接口
// alibaba.latour.strategy.show
//
// 阿里巴巴权益平台权益投放接口
func AlibabaLatourStrategyShow(ctx context.Context, clt *core.SDKClient, req *promotion.AlibabaLatourStrategyShowAPIRequest, resp *promotion.AlibabaLatourStrategyShowAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
