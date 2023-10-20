package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuCombineskuAdd 组合商品新增接口
// alibaba.wdk.sku.combinesku.add
//
// 组合商品新增接口
func AlibabaWdkSkuCombineskuAdd(clt *core.SDKClient, req *wdk.AlibabaWdkSkuCombineskuAddAPIRequest, resp *wdk.AlibabaWdkSkuCombineskuAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
