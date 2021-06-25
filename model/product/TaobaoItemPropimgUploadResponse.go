package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
添加或修改属性图片 APIResponse
taobao.item.propimg.upload

添加一张商品属性图片到num_iid指定的商品中 <br/>传入的num_iid所对应的商品必须属于当前会话的用户 <br/>图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的 <br/>商品属性图片只有享有服务的卖家（如：淘宝大卖家、订购了淘宝多图服务的卖家）才能上传 <br/>商品属性图片有数量和大小上的限制，最多不能超过24张（每个颜色属性都有一张）。
*/
type TaobaoItemPropimgUploadAPIResponse struct {
    model.CommonResponse
    Response *TaobaoItemPropimgUploadResponse `json:"taobao_item_propimg_upload_response,omitempty"`
}

type TaobaoItemPropimgUploadResponse struct {

    // PropImg属性图片结构
    PropImg   *PropImg `json:"prop_img,omitempty"`

}
