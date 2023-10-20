package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Taobaofliggyflightagentauxproductpush 飞猪机票辅营商品投放
// taobao.fliggy.flight.agent.auxproduct.push
//
// 廉航辅营产品投放接口
func Taobaofliggyflightagentauxproductpush(clt *core.SDKClient, req *flight.TaobaofliggyflightagentauxproductpushAPIRequest, session string) (*flight.TaobaofliggyflightagentauxproductpushAPIResponse, error) {
	var resp flight.TaobaofliggyflightagentauxproductpushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
