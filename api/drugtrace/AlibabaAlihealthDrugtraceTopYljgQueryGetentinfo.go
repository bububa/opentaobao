package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopYljgQueryGetentinfo 通过企业名得到唯一标识ref_ent_id及企业ent_id
// alibaba.alihealth.drugtrace.top.yljg.query.getentinfo
//
// 根据企业名称查询ID
func AlibabaAlihealthDrugtraceTopYljgQueryGetentinfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
