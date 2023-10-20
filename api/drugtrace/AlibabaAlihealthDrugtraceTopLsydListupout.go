package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopLsydListupout 零售药店查询本企业上游企业出库单据信息
// alibaba.alihealth.drugtrace.top.lsyd.listupout
//
// 查询货主/本企业上游企业出库单据信息
func AlibabaAlihealthDrugtraceTopLsydListupout(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopLsydListupoutAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
