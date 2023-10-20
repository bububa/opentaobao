package cainiaolocker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaolocker"
)

// Cainiaoendpointlockertopordernoticesendquery 查询订单是否由裹裹发送消息
// cainiao.endpoint.locker.top.order.noticesend.query
//
// 合作公司查询消息发送的接口，判断是否裹裹发送消息
func Cainiaoendpointlockertopordernoticesendquery(clt *core.SDKClient, req *cainiaolocker.CainiaoendpointlockertopordernoticesendqueryAPIRequest, session string) (*cainiaolocker.CainiaoendpointlockertopordernoticesendqueryAPIResponse, error) {
	var resp cainiaolocker.CainiaoendpointlockertopordernoticesendqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
