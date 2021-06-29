package alihealthmedical

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
三方IM图片音频消息上传 APIResponse
alibaba.alihealth.medical.im.data.upload

三方IM图片音频消息上传
*/
type AlibabaAlihealthMedicalImDataUploadAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthMedicalImDataUploadResponse
}

type AlibabaAlihealthMedicalImDataUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_medical_im_data_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
