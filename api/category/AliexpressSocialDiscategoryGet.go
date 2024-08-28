package category

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/category"
)

// AliexpressSocialDiscategoryGet 展示类目获取接口
// aliexpress.social.discategory.get
//
// AE展示类目获取接口
func AliexpressSocialDiscategoryGet(ctx context.Context, clt *core.SDKClient, req *category.AliexpressSocialDiscategoryGetAPIRequest, resp *category.AliexpressSocialDiscategoryGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
