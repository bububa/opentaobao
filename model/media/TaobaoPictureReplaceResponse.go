package media

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
替换图片 APIResponse
taobao.picture.replace

替换一张图片，只替换图片数据，图片名称，图片分类等保持不变。
*/
type TaobaoPictureReplaceAPIResponse struct {
    model.CommonResponse
    TaobaoPictureReplaceResponse
}

type TaobaoPictureReplaceResponse struct {
    XMLName xml.Name `xml:"picture_replace_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 图片替换是否成功
    
    Done   bool `json:"done,omitempty" xml:"done,omitempty"`

    
}
