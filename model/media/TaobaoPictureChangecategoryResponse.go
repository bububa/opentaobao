package media

import (
    "github.com/bububa/opentaobao/model"
)

/* 
修改图片的分类 APIResponse
taobao.picture.changecategory

把批量的图片移动到某个分类下
*/
type TaobaoPictureChangecategoryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPictureChangecategoryResponse `json:"taobao_picture_changecategory_response,omitempty"`
}

type TaobaoPictureChangecategoryResponse struct {

    // 移动图片是否成功：部分移动成功为true，全部移动失败为false。
    Done   bool `json:"done,omitempty"`

}
