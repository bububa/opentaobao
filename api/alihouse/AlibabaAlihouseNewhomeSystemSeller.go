package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeSystemSeller 商品发布授权
// alibaba.alihouse.newhome.system.seller
//
// 商品发布授权
func AlibabaAlihouseNewhomeSystemSeller(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeSystemSellerAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeSystemSellerAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
