package media

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
上传单张图片 APIResponse
taobao.picture.upload

图片空间上传接口
*/
type TaobaoPictureUploadAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"picture_upload_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 当前上传的一张图片信息
    
    Picture   *Picture `json:"picture,omitempty" xml:"