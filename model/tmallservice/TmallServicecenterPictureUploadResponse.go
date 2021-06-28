package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
上传图片 APIResponse
tmall.servicecenter.picture.upload

给服务商ERP系统使用，用于上传图片保存在天猫，一般用于工单信息回传时候保存服务商的服务证明信息相关的图片。
*/
type TmallServicecenterPictureUploadAPIResponse struct {
    model.CommonResponse
    // Response *TmallServicecenterPictureUploadResponse `json:"tmall_servicecenter_picture_upload_response,omitempty"` 
    TmallServicecenterPictureUploadResponse
}

/* model for simplify = false
type TmallServicecenterPictureUploadResponse struct {

    // result
    
    Result  *struct {
        ResultBase  *ResultBase `json:"result_base,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallServicecenterPictureUploadResponse struct {

    // result
    Result   *ResultBase `json:"result,omitempty"`

}
