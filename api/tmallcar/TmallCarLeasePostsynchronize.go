package tmallcar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallCarLeasePostsynchronize 天猫开新车租后信息同步
// tmall.car.lease.postsynchronize
//
// 商家同步天猫开新车租后方案
func TmallCarLeasePostsynchronize(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallCarLeasePostsynchronizeAPIRequest, resp *tmallcar.TmallCarLeasePostsynchronizeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
