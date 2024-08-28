package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallProductSpecsTicketGet 产品规格审核信息获取接口
// tmall.product.specs.ticket.get
//
// 批量根据specId查询产品规格审核信息包括产品规格状态，申请人，拒绝原因等
func TmallProductSpecsTicketGet(ctx context.Context, clt *core.SDKClient, req *product.TmallProductSpecsTicketGetAPIRequest, resp *product.TmallProductSpecsTicketGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
