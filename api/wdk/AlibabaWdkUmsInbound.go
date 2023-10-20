package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkUmsInbound 入库-ERP下发单
// alibaba.wdk.ums.inbound
//
// 入库-ERP下发单
func AlibabaWdkUmsInbound(clt *core.SDKClient, req *wdk.AlibabaWdkUmsInboundAPIRequest, resp *wdk.AlibabaWdkUmsInboundAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
