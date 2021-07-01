package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenStockoutCreateAPIRequest
出库单创建接口 API请求
taobao.qimen.stockout.create

ERP调用奇门接口，创建出库单信息 */
type TaobaoQimenStockoutCreateAPIRequest struct {
	model.Params
	//
	_request *StockOutCreateRequest
}

// New
