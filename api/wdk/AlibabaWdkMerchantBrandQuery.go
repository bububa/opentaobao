package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMerchantBrandQuery 品牌查询接口
// alibaba.wdk.merchant.brand.query
//
// 三江erp对接时，提供品牌查询的接口
func AlibabaWdkMerchantBrandQuery(clt *core.SDKClient, req *wdk.AlibabaWdkMerchantBrandQueryAPIRequest, resp *wdk.AlibabaWdkMerchantBrandQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
