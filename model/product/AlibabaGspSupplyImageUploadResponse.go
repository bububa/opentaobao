package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
gsp图片上传 APIResponse
alibaba.gsp.supply.image.upload

上传图片至目标海外平台的素材空间
*/
type AlibabaGspSupplyImageUploadAPIResponse struct {
    model.CommonResponse
    Response *AlibabaGspSupplyImageUploadResponse `json:"alibaba_gsp_supply_image_upload_response,omitempty"`
}

type AlibabaGspSupplyImageUploadResponse struct {

    // 执行结果
    ServiceSuccess   bool `json:"service_success,omitempty"`

    // 错误码
    ServiceErrorCode   string `json:"service_error_code,omitempty"`

    // 是否重试
    NeedRetry   bool `json:"need_retry,omitempty"`

    // 数据
    Model   *UploadImageResp `json:"model,omitempty"`

    // 错误信息
    ServiceErrorMsg   string `json:"service_error_msg,omitempty"`

}
