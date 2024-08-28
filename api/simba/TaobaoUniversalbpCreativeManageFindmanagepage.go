package simba

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpCreativeManageFindmanagepage 创意库查询创意列表
// taobao.universalbp.creative.manage.findmanagepage
//
// 创意库查询创意列表
func TaobaoUniversalbpCreativeManageFindmanagepage(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpCreativeManageFindmanagepageAPIRequest, resp *simba.TaobaoUniversalbpCreativeManageFindmanagepageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
