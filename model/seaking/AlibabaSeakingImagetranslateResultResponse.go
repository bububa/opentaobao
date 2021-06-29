package seaking

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取图片翻译任务结果 APIResponse
alibaba.seaking.imagetranslate.result

获取图片翻译任务结果
*/
type AlibabaSeakingImagetranslateResultAPIResponse struct {
    model.CommonResponse
    AlibabaSeakingImagetranslateResultResponse
}

type AlibabaSeakingImagetranslateResultResponse struct {
    XMLName xml.Name `xml:"alibaba_seaking_imagetranslate_result_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TaskResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
