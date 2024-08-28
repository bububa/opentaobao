package xhotel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotel"
)

// TaobaoXhotelBnbpromoUpdate 民宿营销活动更新
// taobao.xhotel.bnbpromo.update
//
// 全量更新对应外部活动code相关的营销活动信息
func TaobaoXhotelBnbpromoUpdate(ctx context.Context, clt *core.SDKClient, req *xhotel.TaobaoXhotelBnbpromoUpdateAPIRequest, resp *xhotel.TaobaoXhotelBnbpromoUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
