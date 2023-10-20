package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuCategoryQuery 商家类目获取接口
// alibaba.wdk.sku.category.query
//
// 商家类目获取接口
func AlibabaWdkSkuCategoryQuery(clt *core.SDKClient, req *wdk.AlibabaWdkSkuCategoryQueryAPIRequest, resp *wdk.AlibabaWdkSkuCategoryQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
