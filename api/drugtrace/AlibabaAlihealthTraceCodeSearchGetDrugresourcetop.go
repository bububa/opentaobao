package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaAlihealthTraceCodeSearchGetDrugresourcetop
根据码获取码信息
alibaba.alihealth.trace.code.search.get.drugresourcetop

根据码获取码信息 */
func AlibabaAlihealthTraceCodeSearchGetDrugresourcetop(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIRequest, session string) (*drugtrace.AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthTraceCodeSearchGetDrugresourcetopAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
