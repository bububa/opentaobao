package rhino

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rhino"
)

// TaobaoRhinoCrmGateway crm实体变更回调接口
// taobao.rhino.crm.gateway
//
// crm实体变更回调接口
func TaobaoRhinoCrmGateway(clt *core.SDKClient, req *rhino.TaobaoRhinoCrmGatewayAPIRequest, session string) (*rhino.TaobaoRhinoCrmGatewayAPIResponse, error) {
	var resp rhino.TaobaoRhinoCrmGatewayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
