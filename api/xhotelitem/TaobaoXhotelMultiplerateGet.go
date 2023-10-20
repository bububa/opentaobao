package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelMultiplerateGet 复杂房价查询接口
// taobao.xhotel.multiplerate.get
//
// 查询复杂房价，支持通过入住人数，连住天数，商品信息，房价信息查询
func TaobaoXhotelMultiplerateGet(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelMultiplerateGetAPIRequest, resp *xhotelitem.TaobaoXhotelMultiplerateGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
