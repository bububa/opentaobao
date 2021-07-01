package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusSeatpriceGetAPIRequest
汽车票余票接口 API请求
taobao.bus.seatprice.get

提供给商家，查询汽车票班次余票 */
type TaobaoBusSeatpriceGetAPIRequest struct {
	model.Params
	// 余票请求参数
	_paramBusSeatPriceRQ *BusSeatPriceRq
}

// New
