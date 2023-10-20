package cainiaolocker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaolocker"
)

// Cainiaoendpointlockertoporderwithhold 代扣支付
// cainiao.endpoint.locker.top.order.withhold
//
// 提供代扣，允许有一笔欠款。
func Cainiaoendpointlockertoporderwithhold(clt *core.SDKClient, req *cainiaolocker.CainiaoendpointlockertoporderwithholdAPIRequest, session string) (*cainiaolocker.CainiaoendpointlockertoporderwithholdAPIResponse, error) {
	var resp cainiaolocker.CainiaoendpointlockertoporderwithholdAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
