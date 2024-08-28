package xhotelitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBaseinfoGet 酒店基础信息查询接口
// taobao.xhotel.baseinfo.get
//
// 酒店基础信息(酒店/房型/房价定义)查询接口， 包括 酒店房型可售, 以及 hid 下 的标准房型列表
func TaobaoXhotelBaseinfoGet(ctx context.Context, clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBaseinfoGetAPIRequest, resp *xhotelitem.TaobaoXhotelBaseinfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
