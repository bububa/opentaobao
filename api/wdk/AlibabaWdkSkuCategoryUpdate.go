package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuCategoryUpdate 商家类目修改接口
// alibaba.wdk.sku.category.update
//
// 商家类目修改接口
func AlibabaWdkSkuCategoryUpdate(clt *core.SDKClient, req *wdk.AlibabaWdkSkuCategoryUpdateAPIRequest, resp *wdk.AlibabaWdkSkuCategoryUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
