package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeSystemSeller 商品发布授权
// alibaba.alihouse.newhome.system.seller
//
// 商品发布授权
func AlibabaAlihouseNewhomeSystemSeller(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeSystemSellerAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeSystemSellerAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeSystemSellerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
