package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkBillList 五道口账单拉取接口
// alibaba.wdk.bill.list
//
// 五道口账单拉取接口
func AlibabaWdkBillList(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkBillListAPIRequest, resp *wdk.AlibabaWdkBillListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
