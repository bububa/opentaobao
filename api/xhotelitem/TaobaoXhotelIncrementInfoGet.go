package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelIncrementInfoGet 酒店状态增量查询接口
// taobao.xhotel.increment.info.get
//
// 酒店状态增量查询接口
func TaobaoXhotelIncrementInfoGet(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelIncrementInfoGetAPIRequest, resp *xhotelitem.TaobaoXhotelIncrementInfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
