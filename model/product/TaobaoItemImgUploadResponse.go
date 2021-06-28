package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
添加商品图片 APIResponse
taobao.item.img.upload

添加一张商品图片到num_iid指定的商品中 
传入的num_iid所对应的商品必须属于当前会话的用户 
如果更新图片需要设置itemimg_id，且该itemimg_id的图片记录需要属于传入的num_iid对应的商品。如果新增图片则不用设置 。
使用taobao.item.seller.get中返回的item_imgs字段获取图片id。
商品图片有数量和大小上的限制，根据卖家享有的服务（如：卖家订购了多图服务等），商品图片数量限制不同。
*/
type TaobaoItemImgUploadAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"item_img_upload_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 商品图片结构
    
    ItemImg   *ItemImg `json:"item_img,omitempty" xml:"