package media

import (
    "github.com/bububa/opentaobao/model"
)

/* 
图片总数查询 APIResponse
taobao.picture.pictures.count

图片总数查询
*/
type TaobaoPicturePicturesCountAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPicturePicturesCountResponse `json:"taobao_picture_pictures_count_response,omitempty"`
}

type TaobaoPicturePicturesCountResponse struct {

    // 查询的文件总数
    Totals   int64 `json:"totals,omitempty"`

}
