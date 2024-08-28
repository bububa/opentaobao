package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// CainiaoCntecItemChangeMessage 商品变更消息
// cainiao.cntec.item.change.message
//
// 供货商商品信息变更消息
func CainiaoCntecItemChangeMessage(ctx context.Context, clt *core.SDKClient, req *product.CainiaoCntecItemChangeMessageAPIRequest, resp *product.CainiaoCntecItemChangeMessageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
