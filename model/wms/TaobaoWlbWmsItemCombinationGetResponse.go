package wms

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询组合商品的组合关系 APIResponse
taobao.wlb.wms.item.combination.get

查询组合商品的组合关系
*/
type TaobaoWlbWmsItemCombinationGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbWmsItemCombinationGetResponse `json:"wlb_wms_item_combination_get_response,omitempty"` 
    TaobaoWlbWmsItemCombinationGetResponse
}

/* model for simplify = false
type TaobaoWlbWmsItemCombinationGetResponse struct {

    // 接口返回结果
    
    Result  *struct {
        TaobaoWlbWmsItemCombinationGetResult  *TaobaoWlbWmsItemCombinationGetResult `json:"taobao_wlb_wms_item_combination_get_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoWlbWmsItemCombinationGetResponse struct {

    // 接口返回结果
    Result   *TaobaoWlbWmsItemCombinationGetResult `json:"result,omitempty"`

}
