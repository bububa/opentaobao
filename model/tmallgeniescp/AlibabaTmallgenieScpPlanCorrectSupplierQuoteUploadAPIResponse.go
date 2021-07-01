package tmallgeniescp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商校准后的配额同步 API返回值 
alibaba.tmallgenie.scp.plan.correct.supplier.quote.upload

供应商校准后的配额同步
*/
type AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponse struct {
    model.CommonResponse
    AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponseModel
}

// 供应商校准后的配额同步 成功返回结果
type AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_correct_supplier_quote_upload_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回描述
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
    // 请求唯一ID
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
    // 返回码
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 对象列表
    DataList   []ChannelQuotaDto `json:"data_list,omitempty" xml:"data_list>channel_quota_dto,omitempty"`
}
