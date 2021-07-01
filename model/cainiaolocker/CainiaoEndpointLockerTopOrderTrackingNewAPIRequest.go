package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoEndpointLockerTopOrderTrackingNewAPIRequest
事件回传接口 API请求
cainiao.endpoint.locker.top.order.tracking.new

用于合作公司通知入站、出站信息到菜鸟，共定义了5个操作，1个投件操作，3个取件操作，1个超时提醒。 */
type CainiaoEndpointLockerTopOrderTrackingNewAPIRequest struct {
	model.Params
	// 回传信息
	_trackInfo *CollectTrackingInfo
}

// New
