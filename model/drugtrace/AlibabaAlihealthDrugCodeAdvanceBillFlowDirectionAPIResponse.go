package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单据流向查询 API返回值 
alibaba.alihealth.drug.code.advance.bill.flow.direction

单据流向查询
*/
type AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponseModel
}

// 单据流向查询 成功返回结果
type AlibabaAlihealthDrugCodeAdvanceBillFlowDirectionAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_code_advance_bill_flow_direction_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 和三方交互最外层model对象
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
