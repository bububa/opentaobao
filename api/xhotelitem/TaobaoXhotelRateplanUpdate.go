package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRateplanUpdate 价格计划rateplan更新或添加
// taobao.xhotel.rateplan.update
//
// 酒店产品库rateplan更新或添加
func TaobaoXhotelRateplanUpdate(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateplanUpdateAPIRequest, resp *xhotelitem.TaobaoXhotelRateplanUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
