package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthAlipaypfmOrderSyncAPIRequest
订单数据回传接口 API请求
alibaba.alihealth.alipaypfm.order.sync

订单数据回传接口，各个isv通过我们渠道产生订单需要回传进行统计 */
type AlibabaAlihealthAlipaypfmOrderSyncAPIRequest struct {
	model.Params
	// user_id
	_userId int64
	// 订单id
	_orderId string
	// 订单价格
	_orderPrice string
	// 订单状态
	_orderStatus string
	// 扩展参数
	_extParam string
}

// New
