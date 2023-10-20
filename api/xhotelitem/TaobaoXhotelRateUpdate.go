package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRateUpdate 价格推送接口（全量更新）
// taobao.xhotel.rate.update
//
// 酒店产品库rate更新
func TaobaoXhotelRateUpdate(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRateUpdateAPIRequest, resp *xhotelitem.TaobaoXhotelRateUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
