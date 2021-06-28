package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
上传产品规格认证图片 APIResponse
tmall.product.spec.pic.upload

上传指定类型的产品规格认证文件，并返回存有上传成功图片url的产品规格对象
*/
type TmallProductSpecPicUploadAPIResponse struct {
    model.CommonResponse
    // Response *TmallProductSpecPicUploadResponse `json:"tmall_product_spec_pic_upload_response,omitempty"` 
    TmallProductSpecPicUploadResponse
}

/* model for simplify = false
type TmallProductSpecPicUploadResponse struct {

    // 上传成功的产品规格认证图片url
    
    SpecPicUrl   string `json:"spec_pic_url,omitempty"`
    

}
*/

type TmallProductSpecPicUploadResponse struct {

    // 上传成功的产品规格认证图片url
    SpecPicUrl   string `json:"spec_pic_url,omitempty"`

}
