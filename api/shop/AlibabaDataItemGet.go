package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// Alibabadataitemget 获取商品
// alibaba.data.item.get
//
// 获取商品信息，作为客户端Weex鉴权的虚拟api
func Alibabadataitemget(clt *core.SDKClient, req *shop.AlibabadataitemgetAPIRequest, session string) (*shop.AlibabadataitemgetAPIResponse, error) {
	var resp shop.AlibabadataitemgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
