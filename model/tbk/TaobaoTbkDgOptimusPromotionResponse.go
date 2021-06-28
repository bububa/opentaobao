package tbk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-权益物料精选 APIResponse
taobao.tbk.dg.optimus.promotion

推广者使用。支持入参推广者对应的“推广位”和官方提供的“权益物料id”，获取指定权益物料。
*/
type TaobaoTbkDgOptimusPromotionAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTbkDgOptimusPromotionResponse `json:"tbk_dg_optimus_promotion_response,omitempty"` 
    TaobaoTbkDgOptimusPromotionResponse
}

/* model for simplify = false
type TaobaoTbkDgOptimusPromotionResponse struct {

    // resultList
    
    ResultList  struct {
        TaobaoTbkDgOptimusPromotionMapData  []TaobaoTbkDgOptimusPromotionMapData `json:"taobao_tbk_dg_optimus_promotion_map_data,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

type TaobaoTbkDgOptimusPromotionResponse struct {

    // resultList
    ResultList   []TaobaoTbkDgOptimusPromotionMapData `json:"result_list,omitempty"`

}
