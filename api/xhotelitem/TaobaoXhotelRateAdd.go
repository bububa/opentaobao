package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRateAdd 新增专享房价
// taobao.xhotel.rate.add
//
// 酒店产品库rate添加
func TaobaoXhotelRateAdd(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateAddAPIRequest, resp *xhotelitem.TaobaoXhotelRateAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
