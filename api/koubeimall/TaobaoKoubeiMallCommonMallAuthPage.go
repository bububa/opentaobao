package koubeimall

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/koubeimall"
)

// TaobaoKoubeiMallCommonMallAuthPage 分页查询已授权的商圈列表信息
// taobao.koubei.mall.common.mall.auth.page
//
// 分页查询口碑已授权商圈的列表信息
func TaobaoKoubeiMallCommonMallAuthPage(ctx context.Context, clt *core.SDKClient, req *koubeimall.TaobaoKoubeiMallCommonMallAuthPageAPIRequest, resp *koubeimall.TaobaoKoubeiMallCommonMallAuthPageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
