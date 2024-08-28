package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdGroupCountAdGroup 统计adgroup数量
// alibaba.scbp.ad.group.count.ad.group
//
// 统计adgroup数量
func AlibabaScbpAdGroupCountAdGroup(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdGroupCountAdGroupAPIRequest, resp *scbp.AlibabaScbpAdGroupCountAdGroupAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
