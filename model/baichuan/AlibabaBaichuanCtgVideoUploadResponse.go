package baichuan

import (
    "github.com/bububa/opentaobao/model"
)

/* 
提供优酷的短视频入淘API APIResponse
alibaba.baichuan.ctg.video.upload

提供优酷的短视频入淘API
*/
type AlibabaBaichuanCtgVideoUploadAPIResponse struct {
    model.CommonResponse
    Response *AlibabaBaichuanCtgVideoUploadResponse `json:"alibaba_baichuan_ctg_video_upload_response,omitempty"`
}

type AlibabaBaichuanCtgVideoUploadResponse struct {

    // result
    Result   *CtgResponse `json:"result,omitempty"`

}
