package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaAlihealthDrugtraceTopLsydQueryRelation
单码关联关系查询
alibaba.alihealth.drugtrace.top.lsyd.query.relation

单码关联关系查询 */
func AlibabaAlihealthDrugtraceTopLsydQueryRelation(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
