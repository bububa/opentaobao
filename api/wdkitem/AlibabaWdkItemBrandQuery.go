package wdkitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// AlibabaWdkItemBrandQuery 品牌信息查询
// alibaba.wdk.item.brand.query
//
// 品牌信息查询
func AlibabaWdkItemBrandQuery(clt *core.SDKClient, req *wdkitem.AlibabaWdkItemBrandQueryAPIRequest, resp *wdkitem.AlibabaWdkItemBrandQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
