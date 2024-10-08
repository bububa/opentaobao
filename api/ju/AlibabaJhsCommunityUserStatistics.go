package ju

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ju"
)

// AlibabaJhsCommunityUserStatistics 聚划算社群用户行为上报
// alibaba.jhs.community.user.statistics
//
// 聚划算社群用户行为上报
func AlibabaJhsCommunityUserStatistics(ctx context.Context, clt *core.SDKClient, req *ju.AlibabaJhsCommunityUserStatisticsAPIRequest, resp *ju.AlibabaJhsCommunityUserStatisticsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
