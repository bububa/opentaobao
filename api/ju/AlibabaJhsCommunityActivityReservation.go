package ju

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// Alibabajhscommunityactivityreservation 社群活动预约
// alibaba.jhs.community.activity.reservation
//
// 社群活动预约
func Alibabajhscommunityactivityreservation(clt *core.SDKClient, req *ju.AlibabajhscommunityactivityreservationAPIRequest, session string) (*ju.AlibabajhscommunityactivityreservationAPIResponse, error) {
	var resp ju.AlibabajhscommunityactivityreservationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
