package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
逆向提交 API返回值 
alibaba.wdk.reverse.creatrefund

逆向申请提交
*/
type AlibabaWdkReverseCreatrefundAPIResponse struct {
    model.CommonResponse
    AlibabaWdkReverseCreatrefundAPIResponseModel
}

// 逆向提交 成功返回结果
type AlibabaWdkReverseCreatrefundAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_reverse_creatrefund_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // ReverseResult
    Result   *ReverseResult `json:"result,omitempty" xml:"result,omitempty"`
}
