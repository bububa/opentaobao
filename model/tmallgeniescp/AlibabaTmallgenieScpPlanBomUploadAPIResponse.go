package tmallgeniescp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
计划BOM同步 API返回值 
alibaba.tmallgenie.scp.plan.bom.upload

计划BOM同步
*/
type AlibabaTmallgenieScpPlanBomUploadAPIResponse struct {
    model.CommonResponse
    AlibabaTmallgenieScpPlanBomUploadAPIResponseModel
}

// 计划BOM同步 成功返回结果
type AlibabaTmallgenieScpPlanBomUploadAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_bom_upload_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回描述
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
    // 请求唯一ID
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
    // 返回码
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 对象列表
    DataList   []BomDto `json:"data_list,omitempty" xml:"data_list>bom_dto,omitempty"`
}
