package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除商品图片 APIResponse
taobao.item.img.delete

删除商品图片
*/
type TaobaoItemImgDeleteAPIResponse struct {
    model.CommonResponse
    Response *TaobaoItemImgDeleteResponse `json:"taobao_item_img_delete_response,omitempty"`
}

type TaobaoItemImgDeleteResponse struct {

    // 商品图片结构
    ItemImg   *ItemImg `json:"item_img,omitempty"`

}
