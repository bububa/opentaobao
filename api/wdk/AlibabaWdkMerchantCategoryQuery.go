package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMerchantCategoryQuery 三江erp对接类目查询接口
// alibaba.wdk.merchant.category.query
//
// 三江erp对接类目查询接口
func AlibabaWdkMerchantCategoryQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMerchantCategoryQueryAPIRequest, resp *wdk.AlibabaWdkMerchantCategoryQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
