package cainiaolocker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaolocker"
)

// CainiaoEndpointLockerTopOrderTrackingNew 事件回传接口
// cainiao.endpoint.locker.top.order.tracking.new
//
// 用于合作公司通知入站、出站信息到菜鸟，共定义了5个操作，1个投件操作，3个取件操作，1个超时提醒。
func CainiaoEndpointLockerTopOrderTrackingNew(clt *core.SDKClient, req *cainiaolocker.CainiaoEndpointLockerTopOrderTrackingNewAPIRequest, session string) (*cainiaolocker.CainiaoEndpointLockerTopOrderTrackingNewAPIResponse, error) {
	var resp cainiaolocker.CainiaoEndpointLockerTopOrderTrackingNewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
