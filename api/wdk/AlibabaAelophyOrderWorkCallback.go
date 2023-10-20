package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaAelophyOrderWorkCallback 仓配作业结果回传接口
// alibaba.aelophy.order.work.callback
//
// 仓配作业结果回传接口
func AlibabaAelophyOrderWorkCallback(clt *core.SDKClient, req *wdk.AlibabaAelophyOrderWorkCallbackAPIRequest, resp *wdk.AlibabaAelophyOrderWorkCallbackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
