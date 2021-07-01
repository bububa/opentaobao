package cainiaohandover

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
解决方案询盘 API返回值 
cainiao.global.solution.inquiry

根据交易单号查询可用的解决方案
*/
type CainiaoGlobalSolutionInquiryAPIResponse struct {
    model.CommonResponse
    CainiaoGlobalSolutionInquiryAPIResponseModel
}

// 解决方案询盘 成功返回结果
type CainiaoGlobalSolutionInquiryAPIResponseModel struct {
    XMLName xml.Name `xml:"cainiao_global_solution_inquiry_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // 错误信息
    ErrorInfo   *ErrorInfo `json:"error_info,omitempty" xml:"error_info,omitempty"`
    // 请求结果
    Result   *OpenSolutionInquiryResponse `json:"result,omitempty" xml:"result,omitempty"`
}
