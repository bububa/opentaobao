package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// TaobaoXhotelRatesLiteIncrUpdate 酒店价格库存轻量级增量接口
// taobao.xhotel.rates.lite.incr.update
//
// 多个rate的库存房价开关的增量更新接口
func TaobaoXhotelRatesLiteIncrUpdate(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelRatesLiteIncrUpdateAPIRequest, resp *xhotelitem.TaobaoXhotelRatesLiteIncrUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
