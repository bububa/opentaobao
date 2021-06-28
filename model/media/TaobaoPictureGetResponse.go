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
	RequestId     string         `json:"request_id,omitempty" xml:"picture_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 图片信息列表
    
    Pictures   []Picture `json:"pictures,omitempty" xml:"