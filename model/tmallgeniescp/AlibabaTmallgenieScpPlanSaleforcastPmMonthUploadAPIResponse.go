package tmallgeniescp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
24-销售月预测数量（产管）回传-月度 API返回值 
alibaba.tmallgenie.scp.plan.saleforcast.pm.month.upload

销售月预测数量（产管）回传-月度
*/
type AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIResponse struct {
    model.CommonResponse
    AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIResponseModel
}

// 24-销售月预测数量（产管）回传-月度 成功返回结果
type AlibabaTmallgenieScpPlanSaleforcastPmMonthUploadAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_saleforcast_pm_month_upload_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回描述
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
    // 请求唯一ID
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
    // 返回码
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
