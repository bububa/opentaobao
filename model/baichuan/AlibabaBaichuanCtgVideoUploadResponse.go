package baichuan

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提供优酷的短视频入淘API APIResponse
alibaba.baichuan.ctg.video.upload

提供优酷的短视频入淘API
*/
type AlibabaBaichuanCtgVideoUploadAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_baichuan_ctg_video_upload_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *CtgResponse `json:"result,omitempty" xml:"