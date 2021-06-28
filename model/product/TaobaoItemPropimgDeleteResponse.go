package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
删除属性图片 APIResponse
taobao.item.propimg.delete

删除propimg_id 所指定的商品属性图片 <br/>传入的num_iid所对应的商品必须属于当前会话的用户 <br/>propimg_id对应的属性图片需要属于num_iid对应的商品
*/
type TaobaoItemPropimgDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoItemPropimgDeleteResponse `json:"item_propimg_delete_response,omitempty"` 
    TaobaoItemPropimgDeleteResponse
}

/* model for simplify = false
type TaobaoItemPropimgDeleteResponse struct {

    // 属性图片结构
    
    PropImg  *struct {
        PropImg  *PropImg `json:"prop_img,omitempty"`
    } `json:"prop_img,omitempty"`
    

}
*/

type TaobaoItemPropimgDeleteResponse struct {

    // 属性图片结构
    PropImg   *PropImg `json:"prop_img,omitempty"`

}
