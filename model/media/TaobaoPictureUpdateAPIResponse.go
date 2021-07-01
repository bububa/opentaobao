package media

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改图片名字 API返回值 
taobao.picture.update

修改指定图片的图片名
*/
type TaobaoPictureUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoPictureUpdateAPIResponseModel
}

// 修改图片名字 成功返回结果
type TaobaoPictureUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"picture_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 更新是否成功
    Done   bool `json:"done,omitempty" xml:"done,omitempty"`
}
