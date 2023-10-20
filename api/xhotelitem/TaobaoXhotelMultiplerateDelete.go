package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelMultiplerateDelete 复杂价格删除
// taobao.xhotel.multiplerate.delete
//
// 酒店产品库rate删除
func TaobaoXhotelMultiplerateDelete(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelMultiplerateDeleteAPIRequest, resp *xhotelitem.TaobaoXhotelMultiplerateDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
