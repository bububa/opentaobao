package icbudropshipping

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbudropshipping"
)

// Alibabadropshippingtokencreate 国际站dropshipping 选品token 创建
// alibaba.dropshipping.token.create
//
// 国际站dropshipping 选品token 创建，用于让买家有权限访问我们指定的 商品场馆
func Alibabadropshippingtokencreate(clt *core.SDKClient, req *icbudropshipping.AlibabadropshippingtokencreateAPIRequest, session string) (*icbudropshipping.AlibabadropshippingtokencreateAPIResponse, error) {
	var resp icbudropshipping.AlibabadropshippingtokencreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
