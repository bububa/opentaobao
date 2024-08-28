package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopFlightExceedapplyGet 商旅机票第三方超标审批单搜索接口
// alitrip.btrip.corpop.flight.exceedapply.get
//
// 商旅机票第三方超标审批单搜索接口
func AlitripBtripCorpopFlightExceedapplyGet(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCorpopFlightExceedapplyGetAPIRequest, resp *btrip.AlitripBtripCorpopFlightExceedapplyGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
