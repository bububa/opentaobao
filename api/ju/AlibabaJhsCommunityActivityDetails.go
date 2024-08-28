package ju

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// AlibabaJhsCommunityActivityDetails 社群活动详情
// alibaba.jhs.community.activity.details
//
// 社群活动详情
func AlibabaJhsCommunityActivityDetails(ctx context.Context, clt *core.SDKClient, req *ju.AlibabaJhsCommunityActivityDetailsAPIRequest, resp *ju.AlibabaJhsCommunityActivityDetailsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
