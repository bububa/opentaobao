package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthtracecodesellermilktracetosourceadddata 奶粉溯源-同步数据
// alibaba.alihealth.tracecodeseller.milk.trace.tosource.add.data
//
// 奶粉溯源-同步数据
func Alibabaalihealthtracecodesellermilktracetosourceadddata(clt *core.SDKClient, req *drugtrace.AlibabaalihealthtracecodesellermilktracetosourceadddataAPIRequest, session string) (*drugtrace.AlibabaalihealthtracecodesellermilktracetosourceadddataAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthtracecodesellermilktracetosourceadddataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
