package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopFlightExceedapplyGet 商旅机票第三方超标审批单搜索接口
// alitrip.btrip.corpop.flight.exceedapply.get
//
// 商旅机票第三方超标审批单搜索接口
func AlitripBtripCorpopFlightExceedapplyGet(clt *core.SDKClient, req *btrip.AlitripBtripCorpopFlightExceedapplyGetAPIRequest, resp *btrip.AlitripBtripCorpopFlightExceedapplyGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
