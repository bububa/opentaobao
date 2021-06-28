package tanx

import (
    "github.com/bububa/opentaobao/model"
)

/* 
资质图片上传接口 APIResponse
taobao.tanx.qualification.picture.upload

资质图片上传接口
*/
type TaobaoTanxQualificationPictureUploadAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTanxQualificationPictureUploadResponse `json:"tanx_qualification_picture_upload_response,omitempty"` 
    TaobaoTanxQualificationPictureUploadResponse
}

/* model for simplify = false
type TaobaoTanxQualificationPictureUploadResponse struct {

    // 返回是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 上传后返回的url
    
    Url   string `json:"url,omitempty"`
    

}
*/

type TaobaoTanxQualificationPictureUploadResponse struct {

    // 返回是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 上传后返回的url
    Url   string `json:"url,omitempty"`

}
