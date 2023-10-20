package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthTraceCodeSearchGetDrugresourcetop 根据码获取码信息
// alibaba.alihealth.trace.code.search.get.drugresourcetop
//
// 根据码获取码信息
func AlibabaAlihealthTraceCodeSearchGetDrugresourcetop(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest, resp *drugtrace.AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
