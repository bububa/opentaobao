package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// CainiaoCntecItemChangeMessage 商品变更消息
// cainiao.cntec.item.change.message
//
// 供货商商品信息变更消息
func CainiaoCntecItemChangeMessage(clt *core.SDKClient, req *product.CainiaoCntecItemChangeMessageAPIRequest, resp *product.CainiaoCntecItemChangeMessageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
