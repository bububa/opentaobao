package media

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改图片名字 APIResponse
taobao.picture.update

修改指定图片的图片名
*/
type TaobaoPictureUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoPictureUpdateResponse
}

type TaobaoPictureUpdateResponse struct {
    XMLName xml.Name `xml:"picture_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 更新是否成功
    
    Done   bool `json:"done,omitempty" xml:"done,omitempty"`

    
}
