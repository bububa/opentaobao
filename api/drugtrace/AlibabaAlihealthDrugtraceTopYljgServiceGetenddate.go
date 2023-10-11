package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopYljgServiceGetenddate 获取服务截止日期
// alibaba.alihealth.drugtrace.top.yljg.service.getenddate
//
// 获取企业服务截止时间
func AlibabaAlihealthDrugtraceTopYljgServiceGetenddate(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugtraceTopYljgServiceGetenddateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
