package seaking

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取标题改写任务结果 APIResponse
alibaba.seaking.titlerewrite.result

获取标题改写任务结果
*/
type AlibabaSeakingTitlerewriteResultAPIResponse struct {
    model.CommonResponse
    AlibabaSeakingTitlerewriteResultResponse
}

type AlibabaSeakingTitlerewriteResultResponse struct {
    XMLName xml.Name `xml:"alibaba_seaking_titlerewrite_result_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TaskResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
