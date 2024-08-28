package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderItemTagOperate 全渠道商品打标与去标
// taobao.omniorder.item.tag.operate
//
// 用于对全渠道商品进行打标、去标（门店发货标，门店自提标，前置拆单标）操作。另外还包括增加、删除、修改分单系统，接单系统配置。
func TaobaoOmniorderItemTagOperate(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniorderItemTagOperateAPIRequest, resp *omniorder.TaobaoOmniorderItemTagOperateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
