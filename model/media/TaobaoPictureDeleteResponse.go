package media

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除图片空间图片 APIResponse
taobao.picture.delete

删除图片空间图片
*/
type TaobaoPictureDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoPictureDeleteResponse
}

type TaobaoPictureDeleteResponse struct {
    XMLName xml.Name `xml:"picture_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否删除
    
    Success   string `json:"success,omitempty" xml:"success,omitempty"`

    
}
