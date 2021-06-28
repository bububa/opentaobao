package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
POS商品查询接口 APIResponse
alibaba.mj.oc.getproductbyscancode

此API用于在银泰商场中，POS端扫码获取商品信息
*/
type AlibabaMjOcGetproductbyscancodeAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMjOcGetproductbyscancodeResponse `json:"alibaba_mj_oc_getproductbyscancode_response,omitempty"` 
    AlibabaMjOcGetproductbyscancodeResponse
}

/* model for simplify = false
type AlibabaMjOcGetproductbyscancodeResponse struct {

    // data
    
    ProductList  struct {
        ScanProduct  []ScanProduct `json:"scan_product,omitempty"`
    } `json:"product_list,omitempty"`
    

}
*/

type AlibabaMjOcGetproductbyscancodeResponse struct {

    // data
    ProductList   []ScanProduct `json:"product_list,omitempty"`

}
