package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelBnbpromoGet 民宿查询营销活动
// taobao.xhotel.bnbpromo.get
//
// 民宿查询营销活动
func TaobaoXhotelBnbpromoGet(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelBnbpromoGetAPIRequest, resp *xhotelitem.TaobaoXhotelBnbpromoGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
