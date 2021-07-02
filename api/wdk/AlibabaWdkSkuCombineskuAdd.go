package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuCombineskuAdd 组合商品新增接口
// alibaba.wdk.sku.combinesku.add
//
// 组合商品新增接口
func AlibabaWdkSkuCombineskuAdd(clt *core.SDKClient, req *wdk.AlibabaWdkSkuCombineskuAddAPIRequest, session string) (*wdk.AlibabaWdkSkuCombineskuAddAPIResponse, error) {
	var resp wdk.AlibabaWdkSkuCombineskuAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
