package media

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
图片是否被引用 APIResponse
taobao.picture.isreferenced.get

查询图片是否被引用，被引用返回true，未被引用返回false
*/
type TaobaoPictureIsreferencedGetAPIResponse struct {
    model.CommonResponse
    TaobaoPictureIsreferencedGetResponse
}

type TaobaoPictureIsreferencedGetResponse struct {
    XMLName xml.Name `xml:"picture_isreferenced_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 图片是否被引用
    
    IsReferenced   bool `json:"is_referenced,omitempty" xml:"is_referenced,omitempty"`

    
}
