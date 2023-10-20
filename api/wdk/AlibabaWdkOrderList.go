package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkOrderList 五道口订单拉取
// alibaba.wdk.order.list
//
// 五道口交易订单拉取接口
func AlibabaWdkOrderList(clt *core.SDKClient, req *wdk.AlibabaWdkOrderListAPIRequest, resp *wdk.AlibabaWdkOrderListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
