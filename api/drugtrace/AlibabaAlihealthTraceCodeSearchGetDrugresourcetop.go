package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthtracecodesearchgetdrugresourcetop 根据码获取码信息
// alibaba.alihealth.trace.code.search.get.drugresourcetop
//
// 根据码获取码信息
func Alibabaalihealthtracecodesearchgetdrugresourcetop(clt *core.SDKClient, req *drugtrace.AlibabaalihealthtracecodesearchgetdrugresourcetopAPIRequest, session string) (*drugtrace.AlibabaalihealthtracecodesearchgetdrugresourcetopAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthtracecodesearchgetdrugresourcetopAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
