package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkFinanceOrderBackflow 财务订单回流
// alibaba.wdk.finance.order.backflow
//
// 星巴克拉取财务订单回流数据
func AlibabaWdkFinanceOrderBackflow(clt *core.SDKClient, req *wdk.AlibabaWdkFinanceOrderBackflowAPIRequest, resp *wdk.AlibabaWdkFinanceOrderBackflowAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
