package cntms

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cntms"
)

// CainiaoCntmsLogisticsOrderConsign 菜鸟配商家仓库发货
// cainiao.cntms.logistics.order.consign
//
// 商家包装打印面单结束后，通知菜鸟包裹要发货
func CainiaoCntmsLogisticsOrderConsign(ctx context.Context, clt *core.SDKClient, req *cntms.CainiaoCntmsLogisticsOrderConsignAPIRequest, resp *cntms.CainiaoCntmsLogisticsOrderConsignAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
