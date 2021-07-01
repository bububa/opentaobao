package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWttTradeServiceGetAPIRequest
获取网厅号卡垂直标信息 API请求
taobao.wtt.trade.service.get

查询网厅订单信息 */
type TaobaoWttTradeServiceGetAPIRequest struct {
	model.Params
	// 订单ID
	_bizOrder int64
}

// New
