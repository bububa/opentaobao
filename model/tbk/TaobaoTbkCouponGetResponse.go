package tbk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-公用-阿里妈妈推广券详情查询 APIResponse
taobao.tbk.coupon.get

传入商品ID+券ID(券ID已知情况下)，或者传入me参数，均可查询阿里妈妈推广券详细信息。
*/
type TaobaoTbkCouponGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTbkCouponGetResponse `json:"tbk_coupon_get_response,omitempty"` 
    TaobaoTbkCouponGetResponse
}

/* model for simplify = false
type TaobaoTbkCouponGetResponse struct {

    // data
    
    Data  *struct {
        TaobaoTbkCouponGetMapData  *TaobaoTbkCouponGetMapData `json:"taobao_tbk_coupon_get_map_data,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type TaobaoTbkCouponGetResponse struct {

    // data
    Data   *TaobaoTbkCouponGetMapData `json:"data,omitempty"`

}
