package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkSkuScrollQuery 门店商品批量游标方式查询接口
// alibaba.wdk.sku.scroll.query
//
// 通过游标方式批量获取门店商品信息，包括商品条码，商品名称，价格，会员价等信息。
func AlibabaWdkSkuScrollQuery(clt *core.SDKClient, req *wdk.AlibabaWdkSkuScrollQueryAPIRequest, session string) (*wdk.AlibabaWdkSkuScrollQueryAPIResponse, error) {
	var resp wdk.AlibabaWdkSkuScrollQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
