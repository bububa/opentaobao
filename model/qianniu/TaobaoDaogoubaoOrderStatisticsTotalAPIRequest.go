package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDaogoubaoOrderStatisticsTotalAPIRequest
销售订单总额统计 API请求
taobao.daogoubao.order.statistics.total

对接千牛端数字中心 */
type TaobaoDaogoubaoOrderStatisticsTotalAPIRequest struct {
	model.Params
	// 调试时用的传入id
	_debugId string
	// 需要的字段名
	_field string
}

// New
