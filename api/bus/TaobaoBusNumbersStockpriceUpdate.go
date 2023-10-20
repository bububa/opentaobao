package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusNumbersStockpriceUpdate 汽车票更新价格库存
// taobao.bus.numbers.stockprice.update
//
// 用于汽车票代理商更新价格库存
func TaobaoBusNumbersStockpriceUpdate(clt *core.SDKClient, req *bus.TaobaoBusNumbersStockpriceUpdateAPIRequest, session string) (*bus.TaobaoBusNumbersStockpriceUpdateAPIResponse, error) {
	var resp bus.TaobaoBusNumbersStockpriceUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
