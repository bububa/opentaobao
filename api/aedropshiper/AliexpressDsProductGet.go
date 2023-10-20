package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressDsProductGet 商品信息查询
// aliexpress.ds.product.get
//
// 商品信息查询
func AliexpressDsProductGet(clt *core.SDKClient, req *aedropshiper.AliexpressDsProductGetAPIRequest, session string) (*aedropshiper.AliexpressDsProductGetAPIResponse, error) {
	var resp aedropshiper.AliexpressDsProductGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
