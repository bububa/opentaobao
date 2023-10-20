package ju

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// AlibabaJhsCommunityActivityDetails 社群活动详情
// alibaba.jhs.community.activity.details
//
// 社群活动详情
func AlibabaJhsCommunityActivityDetails(clt *core.SDKClient, req *ju.AlibabaJhsCommunityActivityDetailsAPIRequest, session string) (*ju.AlibabaJhsCommunityActivityDetailsAPIResponse, error) {
	var resp ju.AlibabaJhsCommunityActivityDetailsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
