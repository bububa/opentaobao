package alilabs

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabsTmallgenieAuthRefresh 刷新token
// alibaba.ailabs.tmallgenie.auth.refresh
//
// 通过此接口刷新天猫精灵授权token
func AlibabaAilabsTmallgenieAuthRefresh(ctx context.Context, clt *core.SDKClient, req *alilabs.AlibabaAilabsTmallgenieAuthRefreshAPIRequest, resp *alilabs.AlibabaAilabsTmallgenieAuthRefreshAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
