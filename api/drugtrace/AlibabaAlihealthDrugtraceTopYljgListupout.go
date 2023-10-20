package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopYljgListupout 医疗机构查询本企业上游企业出库单据信息
// alibaba.alihealth.drugtrace.top.yljg.listupout
//
// 查询货主/本企业上游企业出库单据信息
func AlibabaAlihealthDrugtraceTopYljgListupout(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgListupoutAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugtraceTopYljgListupoutAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugtraceTopYljgListupoutAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
