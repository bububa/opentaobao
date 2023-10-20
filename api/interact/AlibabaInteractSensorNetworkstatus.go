package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractsensornetworkstatus 网络状态
// alibaba.interact.sensor.networkstatus
//
// 客户端网络状态
func Alibabainteractsensornetworkstatus(clt *core.SDKClient, req *interact.AlibabainteractsensornetworkstatusAPIRequest, session string) (*interact.AlibabainteractsensornetworkstatusAPIResponse, error) {
	var resp interact.AlibabainteractsensornetworkstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
