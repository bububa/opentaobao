package lstlogistics2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstLogisticsNotraceSendAPIRequest
供应商-异云-无需物流发货 API请求
alibaba.lst.logistics.notrace.send

异地云仓的订单，使用无需物流的发货方式来变更订单发货状态 */
type AlibabaLstLogisticsNotraceSendAPIRequest struct {
	model.Params
	// 入参
	_param *SendDummyOrderParam
}

// New
