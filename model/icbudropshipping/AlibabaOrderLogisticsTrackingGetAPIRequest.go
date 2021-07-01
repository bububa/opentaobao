package icbudropshipping

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaOrderLogisticsTrackingGetAPIRequest
阿里巴巴订单物流轨迹查询 API请求
alibaba.order.logistics.tracking.get

阿里巴巴订单物流轨迹查询 */
type AlibabaOrderLogisticsTrackingGetAPIRequest struct {
	model.Params
	// order id
	_tradeId int64
}

// New
