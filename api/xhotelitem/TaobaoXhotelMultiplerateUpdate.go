package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelMultiplerateUpdate 复杂价格推送接口（全量更新）
// taobao.xhotel.multiplerate.update
//
// 酒店产品库复杂rate（多人价，连住价等）更新
// 同时完全涵盖taobao.xhotel.rate.update的功能
func TaobaoXhotelMultiplerateUpdate(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelMultiplerateUpdateAPIRequest, resp *xhotelitem.TaobaoXhotelMultiplerateUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
