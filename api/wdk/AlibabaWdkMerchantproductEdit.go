package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMerchantproductEdit 商家产品服务-编辑产品
// alibaba.wdk.merchantproduct.edit
//
// 商家产品服务-编辑产品
func AlibabaWdkMerchantproductEdit(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMerchantproductEditAPIRequest, resp *wdk.AlibabaWdkMerchantproductEditAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
