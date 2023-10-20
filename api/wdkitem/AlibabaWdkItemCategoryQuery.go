package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// AlibabaWdkItemCategoryQuery 类目查询接口
// alibaba.wdk.item.category.query
//
// 类目查询接口
func AlibabaWdkItemCategoryQuery(clt *core.SDKClient, req *wdkitem.AlibabaWdkItemCategoryQueryAPIRequest, resp *wdkitem.AlibabaWdkItemCategoryQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
