package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusGuardantGatewayCallback 人卡关系回调
// alibaba.campus.guardant.gateway.callback
//
// 门禁供应商回调平台通知同步结果
func AlibabaCampusGuardantGatewayCallback(clt *core.SDKClient, req *campus.AlibabaCampusGuardantGatewayCallbackAPIRequest, session string) (*campus.AlibabaCampusGuardantGatewayCallbackAPIResponse, error) {
	var resp campus.AlibabaCampusGuardantGatewayCallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
