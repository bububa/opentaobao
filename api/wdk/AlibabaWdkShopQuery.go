package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkShopQuery 门店查询接口
// alibaba.wdk.shop.query
//
// 根据门店code查询门店信息
func AlibabaWdkShopQuery(clt *core.SDKClient, req *wdk.AlibabaWdkShopQueryAPIRequest, resp *wdk.AlibabaWdkShopQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
