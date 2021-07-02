package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytListpartsByagent 物流代货主查往来单位接口
// alibaba.alihealth.drug.kyt.listparts.byagent
//
// 代理企业查询往来单位列表
func AlibabaAlihealthDrugKytListpartsByagent(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytListpartsByagentAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytListpartsByagentAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytListpartsByagentAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
