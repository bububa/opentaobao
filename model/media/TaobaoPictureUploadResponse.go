package media

import (
    "github.com/bububa/opentaobao/model"
)

/* 
上传单张图片 APIResponse
taobao.picture.upload

图片空间上传接口
*/
type TaobaoPictureUploadAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPictureUploadResponse `json:"picture_upload_response,omitempty"` 
    TaobaoPictureUploadResponse
}

/* model for simplify = false
type TaobaoPictureUploadResponse struct {

    // 当前上传的一张图片信息
    
    Picture  *struct {
        Picture  *Picture `json:"picture,omitempty"`
    } `json:"picture,omitempty"`
    

}
*/

type TaobaoPictureUploadResponse struct {

    // 当前上传的一张图片信息
    Picture   *Picture `json:"picture,omitempty"`

}
