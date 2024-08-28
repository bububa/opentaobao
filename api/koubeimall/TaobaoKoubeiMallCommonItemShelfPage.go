package koubeimall

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/koubeimall"
)

// TaobaoKoubeiMallCommonItemShelfPage 门店货架商品列表信息查询
// taobao.koubei.mall.common.item.shelf.page
//
// 查询口碑综合体内门店货架商品列表信息接口
func TaobaoKoubeiMallCommonItemShelfPage(ctx context.Context, clt *core.SDKClient, req *koubeimall.TaobaoKoubeiMallCommonItemShelfPageAPIRequest, resp *koubeimall.TaobaoKoubeiMallCommonItemShelfPageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
