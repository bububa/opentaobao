package media

import (
    "github.com/bububa/opentaobao/model"
)

/* 
多媒体平台文件添加 APIResponse
taobao.media.file.add

用户通过top上传文件到多媒体平台
*/
type TaobaoMediaFileAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoMediaFileAddResponse `json:"media_file_add_response,omitempty"` 
    TaobaoMediaFileAddResponse
}

/* model for simplify = false
type TaobaoMediaFileAddResponse struct {

    // 上传到多媒体平台的文件
    
    File  *struct {
        File  *File `json:"file,omitempty"`
    } `json:"file,omitempty"`
    

}
*/

type TaobaoMediaFileAddResponse struct {

    // 上传到多媒体平台的文件
    File   *File `json:"file,omitempty"`

}
