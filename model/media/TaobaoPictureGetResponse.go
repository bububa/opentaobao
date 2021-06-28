package media

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取图片信息 APIResponse
taobao.picture.get

获取图片信息
*/
type TaobaoPictureGetAPIResponse struct {
    model.CommonResponse
    TaobaoPictureGetResponse
}

type TaobaoPictureGetResponse struct {
    XMLName xml.Name `xml:"picture_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 图片信息列表
    
    Pictures   []Picture `json:"pictures,omitempty" xml:"pictures>picture,omitempty"`
    
    
}
