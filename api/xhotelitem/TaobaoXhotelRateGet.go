package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRateGet 酒店产品库rate查询
// taobao.xhotel.rate.get
//
// 酒店产品库rate查询
func TaobaoXhotelRateGet(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateGetAPIRequest, resp *xhotelitem.TaobaoXhotelRateGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
