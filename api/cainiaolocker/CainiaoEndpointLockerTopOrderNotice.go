package cainiaolocker

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaolocker"
)

// CainiaoEndpointLockerTopOrderNotice 手动触发发短信
// cainiao.endpoint.locker.top.order.notice
//
// 合作公司对订单手动触发短信，有次数限制
func CainiaoEndpointLockerTopOrderNotice(ctx context.Context, clt *core.SDKClient, req *cainiaolocker.CainiaoEndpointLockerTopOrderNoticeAPIRequest, resp *cainiaolocker.CainiaoEndpointLockerTopOrderNoticeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
