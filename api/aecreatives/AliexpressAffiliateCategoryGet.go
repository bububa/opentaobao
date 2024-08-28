package aecreatives

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aecreatives"
)

// AliexpressAffiliateCategoryGet AE流量推广类目信息获取API
// aliexpress.affiliate.category.get
//
// 获取AE流量推广类目的API
func AliexpressAffiliateCategoryGet(ctx context.Context, clt *core.SDKClient, req *aecreatives.AliexpressAffiliateCategoryGetAPIRequest, resp *aecreatives.AliexpressAffiliateCategoryGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
