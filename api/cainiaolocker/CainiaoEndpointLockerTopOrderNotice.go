package cainiaolocker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaolocker"
)

// Cainiaoendpointlockertopordernotice 手动触发发短信
// cainiao.endpoint.locker.top.order.notice
//
// 合作公司对订单手动触发短信，有次数限制
func Cainiaoendpointlockertopordernotice(clt *core.SDKClient, req *cainiaolocker.CainiaoendpointlockertopordernoticeAPIRequest, session string) (*cainiaolocker.CainiaoendpointlockertopordernoticeAPIResponse, error) {
	var resp cainiaolocker.CainiaoendpointlockertopordernoticeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
