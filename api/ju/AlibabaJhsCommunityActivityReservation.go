package ju

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// AlibabaJhsCommunityActivityReservation 社群活动预约
// alibaba.jhs.community.activity.reservation
//
// 社群活动预约
func AlibabaJhsCommunityActivityReservation(clt *core.SDKClient, req *ju.AlibabaJhsCommunityActivityReservationAPIRequest, resp *ju.AlibabaJhsCommunityActivityReservationAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
