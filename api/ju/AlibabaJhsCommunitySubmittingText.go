package ju

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// AlibabaJhsCommunitySubmittingText 聚划算社群动态文案下发接口
// alibaba.jhs.community.submitting.text
//
// 聚划算社群动态文案下发接口
func AlibabaJhsCommunitySubmittingText(clt *core.SDKClient, req *ju.AlibabaJhsCommunitySubmittingTextAPIRequest, session string) (*ju.AlibabaJhsCommunitySubmittingTextAPIResponse, error) {
	var resp ju.AlibabaJhsCommunitySubmittingTextAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
