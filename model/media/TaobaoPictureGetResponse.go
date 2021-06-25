package media

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取图片信息 APIResponse
taobao.picture.get

获取图片信息
*/
type TaobaoPictureGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPictureGetResponse `json:"taobao_picture_get_response,omitempty"`
}

type TaobaoPictureGetResponse struct {

    // 图片信息列表
    Pictures   []Picture `json:"pictures,omitempty"`

}
