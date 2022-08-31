package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkItemCurrentpriceQuery 查询商品当前价格
// alibaba.wdk.item.currentprice.query
//
// 通过渠道店id/sku编码/渠道查询商品当前价格，一次请求商品数量&lt;=20,返回结果key为skuCode
func AlibabaWdkItemCurrentpriceQuery(clt *core.SDKClient, req *wdk.AlibabaWdkItemCurrentpriceQueryAPIRequest, session string) (*wdk.AlibabaWdkItemCurrentpriceQueryAPIResponse, error) {
	var resp wdk.AlibabaWdkItemCurrentpriceQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
