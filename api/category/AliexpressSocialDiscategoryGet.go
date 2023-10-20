package category

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/category"
)

// AliexpressSocialDiscategoryGet 展示类目获取接口
// aliexpress.social.discategory.get
//
// AE展示类目获取接口
func AliexpressSocialDiscategoryGet(clt *core.SDKClient, req *category.AliexpressSocialDiscategoryGetAPIRequest, resp *category.AliexpressSocialDiscategoryGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
