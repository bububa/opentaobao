package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpmaterialaccessallowed 物料准入判断
// taobao.universalbp.material.accessallowed
//
// 入参推广信息，出参相关主体是否可投放。如果投放了风控不准入的商品，无法正常投放。
func Taobaouniversalbpmaterialaccessallowed(clt *core.SDKClient, req *simba.TaobaouniversalbpmaterialaccessallowedAPIRequest, session string) (*simba.TaobaouniversalbpmaterialaccessallowedAPIResponse, error) {
	var resp simba.TaobaouniversalbpmaterialaccessallowedAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
