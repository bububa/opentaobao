package cainiaolocker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaolocker"
)

// CainiaoEndpointLockerTopOrderNotice 手动触发发短信
// cainiao.endpoint.locker.top.order.notice
//
// 合作公司对订单手动触发短信，有次数限制
func CainiaoEndpointLockerTopOrderNotice(clt *core.SDKClient, req *cainiaolocker.CainiaoEndpointLockerTopOrderNoticeAPIRequest, session string) (*cainiaolocker.CainiaoEndpointLockerTopOrderNoticeAPIResponse, error) {
	var resp cainiaolocker.CainiaoEndpointLockerTopOrderNoticeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
