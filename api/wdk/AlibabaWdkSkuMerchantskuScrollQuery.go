package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkSkuMerchantskuScrollQuery
商家商品批量查询接口
alibaba.wdk.sku.merchantsku.scroll.query

提供主档商品数据接口查询 */
func AlibabaWdkSkuMerchantskuScrollQuery(clt *core.SDKClient, req *wdk.AlibabaWdkSkuMerchantskuScrollQueryAPIRequest, session string) (*wdk.AlibabaWdkSkuMerchantskuScrollQueryAPIResponse, error) {
	var resp wdk.AlibabaWdkSkuMerchantskuScrollQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
