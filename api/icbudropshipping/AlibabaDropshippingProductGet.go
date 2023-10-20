package icbudropshipping

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbudropshipping"
)

// Alibabadropshippingproductget 阿里巴巴dropshipping 产品信息获取
// alibaba.dropshipping.product.get
//
// 阿里巴巴dropshipping 产品信息获取
func Alibabadropshippingproductget(clt *core.SDKClient, req *icbudropshipping.AlibabadropshippingproductgetAPIRequest, session string) (*icbudropshipping.AlibabadropshippingproductgetAPIResponse, error) {
	var resp icbudropshipping.AlibabadropshippingproductgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
