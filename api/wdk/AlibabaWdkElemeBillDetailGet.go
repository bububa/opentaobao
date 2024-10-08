package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkElemeBillDetailGet 饿了么对账单查询，带订单明细
// alibaba.wdk.eleme.bill.detail.get
//
// 查询饿了么对账单信息，带订单明细
func AlibabaWdkElemeBillDetailGet(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkElemeBillDetailGetAPIRequest, resp *wdk.AlibabaWdkElemeBillDetailGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
