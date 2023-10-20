package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuQuery 查询商品
// alibaba.wdk.sku.query
//
// 查询商品
func AlibabaWdkSkuQuery(clt *core.SDKClient, req *wdk.AlibabaWdkSkuQueryAPIRequest, resp *wdk.AlibabaWdkSkuQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
