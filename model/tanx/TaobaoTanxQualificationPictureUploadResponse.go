package tanx

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
资质图片上传接口 APIResponse
taobao.tanx.qualification.picture.upload

资质图片上传接口
*/
type TaobaoTanxQualificationPictureUploadAPIResponse struct {
    model.CommonResponse
    TaobaoTanxQualificationPictureUploadResponse
}

type TaobaoTanxQualificationPictureUploadResponse struct {
    XMLName xml.Name `xml:"tanx_qualification_picture_upload_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 上传后返回的url
    
    Url   string `json:"url,omitempty" xml:"url,omitempty"`

    
}
