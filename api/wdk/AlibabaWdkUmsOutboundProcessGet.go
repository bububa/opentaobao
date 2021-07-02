package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkUmsOutboundProcessGet 出库业务UMS异步处理结果返回
// alibaba.wdk.ums.outbound.process.get
//
// 出库业务UMS异步处理结果返回
func AlibabaWdkUmsOutboundProcessGet(clt *core.SDKClient, req *wdk.AlibabaWdkUmsOutboundProcessGetAPIRequest, session string) (*wdk.AlibabaWdkUmsOutboundProcessGetAPIResponse, error) {
	var resp wdk.AlibabaWdkUmsOutboundProcessGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
