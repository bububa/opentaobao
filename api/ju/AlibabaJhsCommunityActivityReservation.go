package ju

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// AlibabaJhsCommunityActivityReservation 社群活动预约
// alibaba.jhs.community.activity.reservation
//
// 社群活动预约
func AlibabaJhsCommunityActivityReservation(clt *core.SDKClient, req *ju.AlibabaJhsCommunityActivityReservationAPIRequest, session string) (*ju.AlibabaJhsCommunityActivityReservationAPIResponse, error) {
	var resp ju.AlibabaJhsCommunityActivityReservationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
