package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据id查询商品 APIResponse
taobao.scitem.get

根据id查询商品
*/
type TaobaoScitemGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoScitemGetResponse `json:"scitem_get_response,omitempty"` 
    TaobaoScitemGetResponse
}

/* model for simplify = false
type TaobaoScitemGetResponse struct {

    // 后端商品
    
    ScItem  *struct {
        ScItem  *ScItem `json:"sc_item,omitempty"`
    } `json:"sc_item,omitempty"`
    

}
*/

type TaobaoScitemGetResponse struct {

    // 后端商品
    ScItem   *ScItem `json:"sc_item,omitempty"`

}
