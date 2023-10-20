package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopLsydQueryRelation 单码关联关系查询
// alibaba.alihealth.drugtrace.top.lsyd.query.relation
//
// 单码关联关系查询
func AlibabaAlihealthDrugtraceTopLsydQueryRelation(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
