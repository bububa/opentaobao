package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelUpdate 酒店更新接口（ID不存在自动新增）
// taobao.xhotel.update
//
// 酒店更新接口
func TaobaoXhotelUpdate(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelUpdateAPIRequest, resp *xhotelitem.TaobaoXhotelUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
