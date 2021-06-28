package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据outerCode查询商品 APIResponse
taobao.scitem.outercode.get

根据outerCode查询商品
*/
type TaobaoScitemOutercodeGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoScitemOutercodeGetResponse `json:"scitem_outercode_get_response,omitempty"` 
    TaobaoScitemOutercodeGetResponse
}

/* model for simplify = false
type TaobaoScitemOutercodeGetResponse struct {

    // 后台商品
    
    ScItem  *struct {
        ScItem  *ScItem `json:"sc_item,omitempty"`
    } `json:"sc_item,omitempty"`
    

}
*/

type TaobaoScitemOutercodeGetResponse struct {

    // 后台商品
    ScItem   *ScItem `json:"sc_item,omitempty"`

}
