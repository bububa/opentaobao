package rhino

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rhino"
)

// TaobaoRhinoCrmGateway crm实体变更回调接口
// taobao.rhino.crm.gateway
//
// crm实体变更回调接口
func TaobaoRhinoCrmGateway(clt *core.SDKClient, req *rhino.TaobaoRhinoCrmGatewayAPIRequest, resp *rhino.TaobaoRhinoCrmGatewayAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
