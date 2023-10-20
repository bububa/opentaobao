package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeAdvanceBillFlowDirection 单据流向查询
// alibaba.alihealth.drug.code.advance.bill.flow.direction
//
// 单据流向查询
func AlibabaAlihealthDrugCodeAdvanceBillFlowDirection(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
