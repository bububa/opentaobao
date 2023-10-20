package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// TaobaoFliggyFlightAgentAuxproductDelete 飞猪机票辅营商品删除
// taobao.fliggy.flight.agent.auxproduct.delete
//
// 廉航辅营产品删除接口
func TaobaoFliggyFlightAgentAuxproductDelete(clt *core.SDKClient, req *flight.TaobaoFliggyFlightAgentAuxproductDeleteAPIRequest, resp *flight.TaobaoFliggyFlightAgentAuxproductDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
