package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripOrderGetAPIRequest
获取欢行统一订单模型 API请求
alibaba.happytrip.order.get

通过订单id获取欢行统一订单模型数据 */
type AlibabaHappytripOrderGetAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
}

// New
