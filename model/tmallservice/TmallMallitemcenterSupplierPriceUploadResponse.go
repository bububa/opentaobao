package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务商服务报价上传 APIResponse
tmall.mallitemcenter.supplier.price.upload

天猫服务商上传服务价格
*/
type TmallMallitemcenterSupplierPriceUploadAPIResponse struct {
    model.CommonResponse
    // Response *TmallMallitemcenterSupplierPriceUploadResponse `json:"tmall_mallitemcenter_supplier_price_upload_response,omitempty"` 
    TmallMallitemcenterSupplierPriceUploadResponse
}

/* model for simplify = false
type TmallMallitemcenterSupplierPriceUploadResponse struct {

    // 接口返回model
    
    Result  *struct {
        TmallMallitemcenterSupplierPriceUploadResult  *TmallMallitemcenterSupplierPriceUploadResult `json:"tmall_mallitemcenter_supplier_price_upload_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallMallitemcenterSupplierPriceUploadResponse struct {

    // 接口返回model
    Result   *TmallMallitemcenterSupplierPriceUploadResult `json:"result,omitempty"`

}
