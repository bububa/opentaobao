package media

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
图片总数查询 APIResponse
taobao.picture.pictures.count

图片总数查询
*/
type TaobaoPicturePicturesCountAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"picture_pictures_count_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询的文件总数
    
    Totals   int64 `json:"totals,omitempty" xml:"