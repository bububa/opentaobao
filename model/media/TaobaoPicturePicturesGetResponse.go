package media

import (
    "github.com/bububa/opentaobao/model"
)

/* 
图片获取 APIResponse
taobao.picture.pictures.get

图片空间对外的图片获取接口，该接口只针对分页获取，获取某一页的图片，该接口不支持总数的查询asd
*/
type TaobaoPicturePicturesGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPicturePicturesGetResponse `json:"taobao_picture_pictures_get_response,omitempty"`
}

type TaobaoPicturePicturesGetResponse struct {

    // 图片空间图片数据对象
    Pictures   []Picture `json:"pictures,omitempty"`

}
