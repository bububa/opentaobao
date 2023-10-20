package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthtracecodesearcgetinfomationvivo 获取vivo banner
// alibaba.alihealth.tracecodesearc.getinfomation.vivo
//
// 获取vivo banner  url
func Alibabaalihealthtracecodesearcgetinfomationvivo(clt *core.SDKClient, req *drugtrace.AlibabaalihealthtracecodesearcgetinfomationvivoAPIRequest, session string) (*drugtrace.AlibabaalihealthtracecodesearcgetinfomationvivoAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthtracecodesearcgetinfomationvivoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
