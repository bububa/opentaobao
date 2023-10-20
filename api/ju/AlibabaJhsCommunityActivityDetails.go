package ju

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// Alibabajhscommunityactivitydetails 社群活动详情
// alibaba.jhs.community.activity.details
//
// 社群活动详情
func Alibabajhscommunityactivitydetails(clt *core.SDKClient, req *ju.AlibabajhscommunityactivitydetailsAPIRequest, session string) (*ju.AlibabajhscommunityactivitydetailsAPIResponse, error) {
	var resp ju.AlibabajhscommunityactivitydetailsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
