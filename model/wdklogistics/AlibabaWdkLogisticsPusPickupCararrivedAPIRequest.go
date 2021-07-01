package wdklogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkLogisticsPusPickupCararrivedAPIRequest
自提业务-车辆到达上报车牌号 API请求
alibaba.wdk.logistics.pus.pickup.cararrived

自提业务-汽车自提,车辆到达上报车牌号 */
type AlibabaWdkLogisticsPusPickupCararrivedAPIRequest struct {
	model.Params
	// 自提点
	_stationCode string
	// 车牌号
	_carNum string
}

// New
