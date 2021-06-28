package media

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
图片获取 APIResponse
taobao.picture.pictures.get

图片空间对外的图片获取接口，该接口只针对分页获取，获取某一页的图片，该接口不支持总数的查询asd
*/
type TaobaoPicturePicturesGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"picture_pictures_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 图片空间图片数据对象
    
    Pictures   []Picture `json:"pictures,omitempty" xml:"