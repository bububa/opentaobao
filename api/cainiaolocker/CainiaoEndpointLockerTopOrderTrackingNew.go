package cainiaolocker

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaolocker"
)

// CainiaoEndpointLockerTopOrderTrackingNew 事件回传接口
// cainiao.endpoint.locker.top.order.tracking.new
//
// 用于合作公司通知入站、出站信息到菜鸟，共定义了5个操作，1个投件操作，3个取件操作，1个超时提醒。
func CainiaoEndpointLockerTopOrderTrackingNew(ctx context.Context, clt *core.SDKClient, req *cainiaolocker.CainiaoEndpointLockerTopOrderTrackingNewAPIRequest, resp *cainiaolocker.CainiaoEndpointLockerTopOrderTrackingNewAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
