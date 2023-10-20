package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelStatusUpdate 酒店状态更新
// taobao.xhotel.status.update
//
// 酒店状态更新
func TaobaoXhotelStatusUpdate(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelStatusUpdateAPIRequest, resp *xhotelitem.TaobaoXhotelStatusUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
