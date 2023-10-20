package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuCategoryAdd 商家类目新增接口
// alibaba.wdk.sku.category.add
//
// 商家类目新增接口
func AlibabaWdkSkuCategoryAdd(clt *core.SDKClient, req *wdk.AlibabaWdkSkuCategoryAddAPIRequest, resp *wdk.AlibabaWdkSkuCategoryAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
