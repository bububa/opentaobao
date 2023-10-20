package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopLsydQueryListparts 往来单位查询
// alibaba.alihealth.drugtrace.top.lsyd.query.listparts
//
// 查询往来单位列表
func AlibabaAlihealthDrugtraceTopLsydQueryListparts(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
