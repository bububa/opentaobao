package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopLsydListupoutDetail 上游出库单单据明细查询
// alibaba.alihealth.drugtrace.top.lsyd.listupout.detail
//
// 查询上游出库单明细(带追溯码信息)
func AlibabaAlihealthDrugtraceTopLsydListupoutDetail(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
