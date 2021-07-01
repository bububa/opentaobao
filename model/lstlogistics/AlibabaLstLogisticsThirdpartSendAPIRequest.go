package lstlogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstLogisticsThirdpartSendAPIRequest
供应商-异云-使用第三方物流发货 API请求
alibaba.lst.logistics.thirdpart.send

异地云仓的订单，使用第三方物流的发货方式来变更订单发货状态 */
type AlibabaLstLogisticsThirdpartSendAPIRequest struct {
	model.Params
	// 入参
	_param *SendOfflineOrderParam
}

// New
