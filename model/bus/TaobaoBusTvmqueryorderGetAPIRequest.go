package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusTvmqueryorderGetAPIRequest
线下自助机查询订单信息 API请求
taobao.bus.tvmqueryorder.get

查询订单详情 */
type TaobaoBusTvmqueryorderGetAPIRequest struct {
	model.Params
	// 阿里订单标编号
	_alitripOrderId string
}

// New
