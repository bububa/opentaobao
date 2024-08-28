package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoRegionPriceManage 编辑区域价格
// taobao.region.price.manage
//
// 编辑区域价格
func TaobaoRegionPriceManage(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoRegionPriceManageAPIRequest, resp *fenxiao.TaobaoRegionPriceManageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
