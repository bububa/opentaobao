package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusNumbersStockpriceUpdateAPIRequest
汽车票更新价格库存 API请求
taobao.bus.numbers.stockprice.update

用于汽车票代理商更新价格库存 */
type TaobaoBusNumbersStockpriceUpdateAPIRequest struct {
	model.Params
	// 请求参数
	_paramTopBusPriceAndStockUpdateRQ *TopBusPriceAndStockUpdateRq
}

// New
