package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuWarehouseskuScrollQuery 仓商品遍历接口(游标)
// alibaba.wdk.sku.warehousesku.scroll.query
//
// 提供仓商品数据接口查询
func AlibabaWdkSkuWarehouseskuScrollQuery(clt *core.SDKClient, req *wdk.AlibabaWdkSkuWarehouseskuScrollQueryAPIRequest, resp *wdk.AlibabaWdkSkuWarehouseskuScrollQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
