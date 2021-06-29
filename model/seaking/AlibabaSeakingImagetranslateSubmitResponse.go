package seaking

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交图片翻译任务 APIResponse
alibaba.seaking.imagetranslate.submit

提交图片翻译任务
*/
type AlibabaSeakingImagetranslateSubmitAPIResponse struct {
    model.CommonResponse
    AlibabaSeakingImagetranslateSubmitResponse
}

type AlibabaSeakingImagetranslateSubmitResponse struct {
    XMLName xml.Name `xml:"alibaba_seaking_imagetranslate_submit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 任务id
    
    TaskId   int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`

    
}
