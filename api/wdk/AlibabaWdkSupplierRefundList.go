package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSupplierRefundList 五道口按供应商拉取退款单
// alibaba.wdk.supplier.refund.list
//
// 五道口按供应商拉取退款单
func AlibabaWdkSupplierRefundList(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkSupplierRefundListAPIRequest, resp *wdk.AlibabaWdkSupplierRefundListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
