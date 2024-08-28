package koubeimall

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/koubeimall"
)

// TaobaoKoubeiMallCommonStorePage 分页查询综合体内的门店列表信息
// taobao.koubei.mall.common.store.page
//
// 分页查询综合体内的门店列表信息
func TaobaoKoubeiMallCommonStorePage(ctx context.Context, clt *core.SDKClient, req *koubeimall.TaobaoKoubeiMallCommonStorePageAPIRequest, resp *koubeimall.TaobaoKoubeiMallCommonStorePageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
