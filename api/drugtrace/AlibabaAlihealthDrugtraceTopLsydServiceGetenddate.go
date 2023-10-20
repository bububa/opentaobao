package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopLsydServiceGetenddate 获取企业服务截止时间
// alibaba.alihealth.drugtrace.top.lsyd.service.getenddate
//
// 获取企业服务截止时间
func AlibabaAlihealthDrugtraceTopLsydServiceGetenddate(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydServiceGetenddateAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugtraceTopLsydServiceGetenddateAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugtraceTopLsydServiceGetenddateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
