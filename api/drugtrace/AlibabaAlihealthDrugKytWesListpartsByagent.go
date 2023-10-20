package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesListpartsByagent 物流代货主查找往来单位接口
// alibaba.alihealth.drug.kyt.wes.listparts.byagent
//
// 代理企业查询往来单位列表
func AlibabaAlihealthDrugKytWesListpartsByagent(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesListpartsByagentAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytWesListpartsByagentAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytWesListpartsByagentAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
