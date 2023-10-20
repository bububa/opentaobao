package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMerchantStoreitemQuery 门店商品信心查询
// alibaba.wdk.merchant.storeitem.query
//
// 门店商品信心查询
func AlibabaWdkMerchantStoreitemQuery(clt *core.SDKClient, req *wdk.AlibabaWdkMerchantStoreitemQueryAPIRequest, resp *wdk.AlibabaWdkMerchantStoreitemQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
