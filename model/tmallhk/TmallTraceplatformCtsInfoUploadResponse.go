package tmallhk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
CTS提交溯源信息 APIResponse
tmall.traceplatform.cts.info.upload

cts上传溯源信息
*/
type TmallTraceplatformCtsInfoUploadAPIResponse struct {
    model.CommonResponse
    Response *TmallTraceplatformCtsInfoUploadResponse `json:"tmall_traceplatform_cts_info_upload_response,omitempty"`
}

type TmallTraceplatformCtsInfoUploadResponse struct {

    // result
    Result   *DataResult `json:"result,omitempty"`

}
