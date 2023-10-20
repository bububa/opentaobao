package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuStoreskuScrollQuery 门店商品批量查询接口
// alibaba.wdk.sku.storesku.scroll.query
//
// 提供门店商品批量查询接口
func AlibabaWdkSkuStoreskuScrollQuery(clt *core.SDKClient, req *wdk.AlibabaWdkSkuStoreskuScrollQueryAPIRequest, resp *wdk.AlibabaWdkSkuStoreskuScrollQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
