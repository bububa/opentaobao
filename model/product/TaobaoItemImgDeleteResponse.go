package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除商品图片 APIResponse
taobao.item.img.delete

删除商品图片
*/
type TaobaoItemImgDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"item_img_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品图片结构
    
    ItemImg   *ItemImg `json:"item_img,omitempty" xml:"