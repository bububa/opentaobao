package tmallgeniescp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
10-同步库存现有量 APIResponse
alibaba.tmallgenie.scp.plan.inventor.qty.get

同步库存现有量
*/
type AlibabaTmallgenieScpPlanInventorQtyGetAPIResponse struct {
    model.CommonResponse
    AlibabaTmallgenieScpPlanInventorQtyGetResponse
}

type AlibabaTmallgenieScpPlanInventorQtyGetResponse struct {
    XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_inventor_qty_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象封装
    
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
