package dmp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dmp"
)

// TaobaoDmpCrowdTemplateApply 人群模版采纳并生成人群API
// taobao.dmp.crowd.template.apply
//
// 人群模版采纳并生成人群API
func TaobaoDmpCrowdTemplateApply(ctx context.Context, clt *core.SDKClient, req *dmp.TaobaoDmpCrowdTemplateApplyAPIRequest, resp *dmp.TaobaoDmpCrowdTemplateApplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
