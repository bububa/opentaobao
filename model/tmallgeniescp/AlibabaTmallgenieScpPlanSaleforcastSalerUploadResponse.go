package tmallgeniescp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
19-销售预测数量（销管）回传接口 APIResponse
alibaba.tmallgenie.scp.plan.saleforcast.saler.upload

销售预测数量（销管）回传接口
*/
type AlibabaTmallgenieScpPlanSaleforcastSalerUploadAPIResponse struct {
    model.CommonResponse
    AlibabaTmallgenieScpPlanSaleforcastSalerUploadResponse
}

type AlibabaTmallgenieScpPlanSaleforcastSalerUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_saleforcast_saler_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 参数msg
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // 请求唯一id
    
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`

    
    // 参数code
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
}
