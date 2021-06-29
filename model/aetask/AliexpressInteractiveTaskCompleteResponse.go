package aetask

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
任务完成接口 APIResponse
aliexpress.interactive.task.complete

用户完成任务
*/
type AliexpressInteractiveTaskCompleteAPIResponse struct {
    model.CommonResponse
    AliexpressInteractiveTaskCompleteResponse
}

type AliexpressInteractiveTaskCompleteResponse struct {
    XMLName xml.Name `xml:"aliexpress_interactive_task_complete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AliexpressInteractiveTaskCompleteResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
