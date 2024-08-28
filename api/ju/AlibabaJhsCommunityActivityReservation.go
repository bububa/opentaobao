package ju

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// AlibabaJhsCommunityActivityReservation 社群活动预约
// alibaba.jhs.community.activity.reservation
//
// 社群活动预约
func AlibabaJhsCommunityActivityReservation(ctx context.Context, clt *core.SDKClient, req *ju.AlibabaJhsCommunityActivityReservationAPIRequest, resp *ju.AlibabaJhsCommunityActivityReservationAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
