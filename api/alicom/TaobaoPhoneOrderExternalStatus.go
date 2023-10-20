package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Taobaophoneorderexternalstatus 话费外放订单状态接口
// taobao.phone.order.external.status
//
// 话费外放订单状态接口
func Taobaophoneorderexternalstatus(clt *core.SDKClient, req *alicom.TaobaophoneorderexternalstatusAPIRequest, session string) (*alicom.TaobaophoneorderexternalstatusAPIResponse, error) {
	var resp alicom.TaobaophoneorderexternalstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
