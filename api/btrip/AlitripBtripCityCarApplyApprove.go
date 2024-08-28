package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCityCarApplyApprove 三方市内用车申请单审批
// alitrip.btrip.city.car.apply.approve
//
// 三方市内用车申请单审批
func AlitripBtripCityCarApplyApprove(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCityCarApplyApproveAPIRequest, resp *btrip.AlitripBtripCityCarApplyApproveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
