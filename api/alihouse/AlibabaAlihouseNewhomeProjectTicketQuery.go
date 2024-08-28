package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeProjectTicketQuery 根据商品id查询核销卷信息
// alibaba.alihouse.newhome.project.ticket.query
//
// 根据商品id查询核销卷信息
func AlibabaAlihouseNewhomeProjectTicketQuery(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeProjectTicketQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
