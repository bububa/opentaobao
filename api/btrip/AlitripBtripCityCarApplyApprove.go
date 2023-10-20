package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCityCarApplyApprove 三方市内用车申请单审批
// alitrip.btrip.city.car.apply.approve
//
// 三方市内用车申请单审批
func AlitripBtripCityCarApplyApprove(clt *core.SDKClient, req *btrip.AlitripBtripCityCarApplyApproveAPIRequest, resp *btrip.AlitripBtripCityCarApplyApproveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
