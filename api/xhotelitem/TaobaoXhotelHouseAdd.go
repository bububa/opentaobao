package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelHouseAdd 非标准民宿房源添加
// taobao.xhotel.house.add
//
// 添加酒店或更新酒店
func TaobaoXhotelHouseAdd(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelHouseAddAPIRequest, resp *xhotelitem.TaobaoXhotelHouseAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
