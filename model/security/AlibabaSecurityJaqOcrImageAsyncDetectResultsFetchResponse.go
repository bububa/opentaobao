package security

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚安全获取异步图文识别结果接口 APIResponse
alibaba.security.jaq.ocr.image.async.detect.results.fetch

获取异步图像字符识别结果接口根据图像检测接口返回taskid来获取对应图像的检测结果
*/
type AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchResponse `json:"alibaba_security_jaq_ocr_image_async_detect_results_fetch_response,omitempty"` 
    AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchResponse
}

/* model for simplify = false
type AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchResponse struct {

    // 出参结构体
    
    Data  *struct {
        JaqImageDetectResultCollection  *JaqImageDetectResultCollection `json:"jaq_image_detect_result_collection,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type AlibabaSecurityJaqOcrImageAsyncDetectResultsFetchResponse struct {

    // 出参结构体
    Data   *JaqImageDetectResultCollection `json:"data,omitempty"`

}
