package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
上传单张产品属性图片，如果需要传多张，可调多次 APIResponse
taobao.product.propimg.upload

传入产品ID <br/>传入props,目前仅支持颜色属性.调用taobao.itemprops.get.v2取得颜色属性pid,<br/>再用taobao.itempropvalues.get取得vid;格式:pid:vid,只能传入一个颜色pid:vid串; <br/>传入图片内容 <br/>注意：图片最大为2M,只支持JPG,GIF,如果需要传多张，可调多次
*/
type TaobaoProductPropimgUploadAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoProductPropimgUploadResponse `json:"product_propimg_upload_response,omitempty"` 
    TaobaoProductPropimgUploadResponse
}

/* model for simplify = false
type TaobaoProductPropimgUploadResponse struct {

    // 支持返回产品属性图片中的：url,id,created,modified
    
    ProductPropImg  *struct {
        ProductPropImg  *ProductPropImg `json:"product_prop_img,omitempty"`
    } `json:"product_prop_img,omitempty"`
    

}
*/

type TaobaoProductPropimgUploadResponse struct {

    // 支持返回产品属性图片中的：url,id,created,modified
    ProductPropImg   *ProductPropImg `json:"product_prop_img,omitempty"`

}
