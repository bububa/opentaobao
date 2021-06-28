package tmallhk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
AWDC提交溯源信息 APIResponse
tmall.traceplatform.awdc.info.upload

天猫溯源-AWDC-上传溯源信息
*/
type TmallTraceplatformAwdcInfoUploadAPIResponse struct {
    model.CommonResponse
    // Response *TmallTraceplatformAwdcInfoUploadResponse `json:"tmall_traceplatform_awdc_info_upload_response,omitempty"` 
    TmallTraceplatformAwdcInfoUploadResponse
}

/* model for simplify = false
type TmallTraceplatformAwdcInfoUploadResponse struct {

    // result
    
    Result  *struct {
        DataResult  *DataResult `json:"data_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallTraceplatformAwdcInfoUploadResponse struct {

    // result
    Result   *DataResult `json:"result,omitempty"`

}
