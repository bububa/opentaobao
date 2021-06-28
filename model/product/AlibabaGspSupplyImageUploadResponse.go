package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
gsp图片上传 APIResponse
alibaba.gsp.supply.image.upload

上传图片至目标海外平台的素材空间
*/
type AlibabaGspSupplyImageUploadAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_gsp_supply_image_upload_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 执行结果
    
    ServiceSuccess   bool `json:"service_success,omitempty" xml:"