package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelServicetimeGet 查询实体对应的服务时间数据
// taobao.xhotel.servicetime.get
//
// 通过实体来获取对应的服务时间数据
func TaobaoXhotelServicetimeGet(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelServicetimeGetAPIRequest, resp *xhotelitem.TaobaoXhotelServicetimeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
