package tmallgeniescp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
同步供应商校准后的配额-二级物料 API返回值 
alibaba.tmallgenie.scp.plan.correct.supplier.quote.raw.upload

同步供应商校准后的配额-二级物料
*/
type AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadAPIResponse struct {
    model.CommonResponse
    AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadResponse
}

// 同步供应商校准后的配额-二级物料 成功返回结果
type AlibabaTmallgenieScpPlanCorrectSupplierQuoteRawUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_correct_supplier_quote_raw_upload_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回描述
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
    // 请求唯一ID
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
    // 返回码
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 对象列表
    DataList   []CurrentQuotaDTO `json:"data_list,omitempty" xml:"data_list>current_quota_dto,omitempty"`
}
