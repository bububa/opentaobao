package icbu

import (
    "github.com/bububa/opentaobao/model"
)

/* 
图片银行图片上传开放接口 APIResponse
alibaba.icbu.photobank.upload

图片银行图片上传开放接口
*/
type AlibabaIcbuPhotobankUploadAPIResponse struct {
    model.CommonResponse
    Response *AlibabaIcbuPhotobankUploadResponse `json:"alibaba_icbu_photobank_upload_response,omitempty"`
}

type AlibabaIcbuPhotobankUploadResponse struct {

    // 图片信息
    UploadImageResponse   *UploadImageResponseDo `json:"upload_image_response,omitempty"`

}
