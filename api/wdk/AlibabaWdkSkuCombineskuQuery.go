package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuCombineskuQuery 组合商品查询接口
// alibaba.wdk.sku.combinesku.query
//
// 查询组合商品接口
func AlibabaWdkSkuCombineskuQuery(clt *core.SDKClient, req *wdk.AlibabaWdkSkuCombineskuQueryAPIRequest, session string) (*wdk.AlibabaWdkSkuCombineskuQueryAPIResponse, error) {
	var resp wdk.AlibabaWdkSkuCombineskuQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
