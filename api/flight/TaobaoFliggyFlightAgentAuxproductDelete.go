package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Taobaofliggyflightagentauxproductdelete 飞猪机票辅营商品删除
// taobao.fliggy.flight.agent.auxproduct.delete
//
// 廉航辅营产品删除接口
func Taobaofliggyflightagentauxproductdelete(clt *core.SDKClient, req *flight.TaobaofliggyflightagentauxproductdeleteAPIRequest, session string) (*flight.TaobaofliggyflightagentauxproductdeleteAPIResponse, error) {
	var resp flight.TaobaofliggyflightagentauxproductdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
