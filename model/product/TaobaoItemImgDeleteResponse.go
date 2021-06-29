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
    TaobaoItemImgDeleteResponse
}

type TaobaoItemImgDeleteResponse struct {
    XMLName xml.Name `xml:"item_img_delete_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 商品图片结构
    
    ItemImg   *ItemImg `json:"item_img,omitempty" xml:"item_img,omitempty"`

    
}
