package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripCityCarApplyAdd 三方市内用车申请单同步
// alitrip.btrip.city.car.apply.add
//
// 三方市内用车申请单同步
func AlitripBtripCityCarApplyAdd(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCityCarApplyAddAPIRequest, resp *btrip.AlitripBtripCityCarApplyAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
