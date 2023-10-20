package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkUmsOrderGet 查询店仓作业单据清单 （库存对账辅助）-回流单
// alibaba.wdk.ums.order.get
//
// 查询店仓作业单据清单 （库存对账辅助）-回流单
func AlibabaWdkUmsOrderGet(clt *core.SDKClient, req *wdk.AlibabaWdkUmsOrderGetAPIRequest, resp *wdk.AlibabaWdkUmsOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
