package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
上传单张产品非主图，如果需要传多张，可调多次 APIResponse
taobao.product.img.upload

1.传入产品ID <br/>2.传入图片内容 <br/>注意：图片最大为500K,只支持JPG,GIF格式,如果需要传多张，可调多次
*/
type TaobaoProductImgUploadAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoProductImgUploadResponse `json:"product_img_upload_response,omitempty"` 
    TaobaoProductImgUploadResponse
}

/* model for simplify = false
type TaobaoProductImgUploadResponse struct {

    // 返回产品图片结构中的：url,id,created,modified
    
    ProductImg  *struct {
        ProductImg  *ProductImg `json:"product_img,omitempty"`
    } `json:"product_img,omitempty"`
    

}
*/

type TaobaoProductImgUploadResponse struct {

    // 返回产品图片结构中的：url,id,created,modified
    ProductImg   *ProductImg `json:"product_img,omitempty"`

}
