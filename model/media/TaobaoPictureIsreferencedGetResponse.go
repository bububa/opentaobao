package media

import (
    "github.com/bububa/opentaobao/model"
)

/* 
图片是否被引用 APIResponse
taobao.picture.isreferenced.get

查询图片是否被引用，被引用返回true，未被引用返回false
*/
type TaobaoPictureIsreferencedGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPictureIsreferencedGetResponse `json:"taobao_picture_isreferenced_get_response,omitempty"`
}

type TaobaoPictureIsreferencedGetResponse struct {

    // 图片是否被引用
    IsReferenced   bool `json:"is_referenced,omitempty"`

}
