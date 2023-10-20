package icbudropshipping

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbudropshipping"
)

// Alibababuynowordercreate 阿里巴巴买家buynow下单接口
// alibaba.buynow.order.create
//
// 阿里巴巴买家下单接口
func Alibababuynowordercreate(clt *core.SDKClient, req *icbudropshipping.AlibababuynowordercreateAPIRequest, session string) (*icbudropshipping.AlibababuynowordercreateAPIResponse, error) {
	var resp icbudropshipping.AlibababuynowordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
