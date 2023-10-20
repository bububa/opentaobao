package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkItemCategoryUpdate 修改类目
// alibaba.wdk.item.category.update
//
// 修改类目
func AlibabaWdkItemCategoryUpdate(clt *core.SDKClient, req *wdk.AlibabaWdkItemCategoryUpdateAPIRequest, resp *wdk.AlibabaWdkItemCategoryUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
