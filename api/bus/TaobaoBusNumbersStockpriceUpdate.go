package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobusnumbersstockpriceupdate 汽车票更新价格库存
// taobao.bus.numbers.stockprice.update
//
// 用于汽车票代理商更新价格库存
func Taobaobusnumbersstockpriceupdate(clt *core.SDKClient, req *bus.TaobaobusnumbersstockpriceupdateAPIRequest, session string) (*bus.TaobaobusnumbersstockpriceupdateAPIResponse, error) {
	var resp bus.TaobaobusnumbersstockpriceupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
