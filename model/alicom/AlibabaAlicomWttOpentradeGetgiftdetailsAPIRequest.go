package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlicomWttOpentradeGetgiftdetailsAPIRequest
存送业务查询奖品信息 API请求
alibaba.alicom.wtt.opentrade.getgiftdetails

话费宝充值送查询奖品信息 */
type AlibabaAlicomWttOpentradeGetgiftdetailsAPIRequest struct {
	model.Params
	// 活动ID
	_activityId string
}

// New
