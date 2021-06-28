package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家查询物流商产品类型接口 APIResponse
cainiao.waybill.ii.product

商家可以查询物流商的产品类型和服务能力。
*/
type CainiaoWaybillIiProductAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoWaybillIiProductResponse `json:"cainiao_waybill_ii_product_response,omitempty"` 
    CainiaoWaybillIiProductResponse
}

/* model for simplify = false
type CainiaoWaybillIiProductResponse struct {

    // 返回值
    
    ProductTypes  struct {
        WaybillProductType  []WaybillProductType `json:"waybill_product_type,omitempty"`
    } `json:"product_types,omitempty"`
    

}
*/

type CainiaoWaybillIiProductResponse struct {

    // 返回值
    ProductTypes   []WaybillProductType `json:"product_types,omitempty"`

}
