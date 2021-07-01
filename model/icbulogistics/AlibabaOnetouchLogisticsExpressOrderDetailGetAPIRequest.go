package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest
订单详细信息(面单及仓库信息) API请求
alibaba.onetouch.logistics.express.order.detail.get

订单详细信息(面单及仓库信息) */
type AlibabaOnetouchLogisticsExpressOrderDetailGetAPIRequest struct {
	model.Params
	// 请求参数
	_paramQuery *LogisticsOrderQueryDto
}

// New
