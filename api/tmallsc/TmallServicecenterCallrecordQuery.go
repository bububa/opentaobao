package tmallsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterCallrecordQuery 天猫服务平台服务商查询通话记录接口
// tmall.servicecenter.callrecord.query
//
// 天猫服务平台服务商查询通话记录接口
func TmallServicecenterCallrecordQuery(ctx context.Context, clt *core.SDKClient, req *tmallsc.TmallServicecenterCallrecordQueryAPIRequest, resp *tmallsc.TmallServicecenterCallrecordQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
