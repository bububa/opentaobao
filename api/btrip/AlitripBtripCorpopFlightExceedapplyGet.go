package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCorpopFlightExceedapplyGet 商旅机票第三方超标审批单搜索接口
// alitrip.btrip.corpop.flight.exceedapply.get
//
// 商旅机票第三方超标审批单搜索接口
func AlitripBtripCorpopFlightExceedapplyGet(clt *core.SDKClient, req *btrip.AlitripBtripCorpopFlightExceedapplyGetAPIRequest, session string) (*btrip.AlitripBtripCorpopFlightExceedapplyGetAPIResponse, error) {
	var resp btrip.AlitripBtripCorpopFlightExceedapplyGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
