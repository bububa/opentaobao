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
    Response *TaobaoPictureUploadResponse `json:"taobao_picture_upload_response,omitempty"`
}

type TaobaoPictureUploadResponse struct {

    // 当前上传的一张图片信息
    Picture   *Picture `json:"picture,omitempty"`

}
