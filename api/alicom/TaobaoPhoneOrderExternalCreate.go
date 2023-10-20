package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Taobaophoneorderexternalcreate 数字虚拟话费外放下单接口
// taobao.phone.order.external.create
//
// 数字虚拟话费外放下单接口
func Taobaophoneorderexternalcreate(clt *core.SDKClient, req *alicom.TaobaophoneorderexternalcreateAPIRequest, session string) (*alicom.TaobaophoneorderexternalcreateAPIResponse, error) {
	var resp alicom.TaobaophoneorderexternalcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
