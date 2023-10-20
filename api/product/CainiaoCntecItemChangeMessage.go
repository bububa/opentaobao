package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Cainiaocntecitemchangemessage 商品变更消息
// cainiao.cntec.item.change.message
//
// 供货商商品信息变更消息
func Cainiaocntecitemchangemessage(clt *core.SDKClient, req *product.CainiaocntecitemchangemessageAPIRequest, session string) (*product.CainiaocntecitemchangemessageAPIResponse, error) {
	var resp product.CainiaocntecitemchangemessageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
