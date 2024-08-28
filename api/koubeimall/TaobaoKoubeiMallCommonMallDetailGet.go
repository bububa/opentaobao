package koubeimall

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/koubeimall"
)

// TaobaoKoubeiMallCommonMallDetailGet 查询商圈详细信息
// taobao.koubei.mall.common.mall.detail.get
//
// 查询口碑综合体-商圈详细信息，包含商圈基础信息、门店类目分类、商圈推荐商品等模块信息
func TaobaoKoubeiMallCommonMallDetailGet(ctx context.Context, clt *core.SDKClient, req *koubeimall.TaobaoKoubeiMallCommonMallDetailGetAPIRequest, resp *koubeimall.TaobaoKoubeiMallCommonMallDetailGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
