package category

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询商品类目属性变更 APIResponse
taobao.item.catprops.modification.get

查询商品类目属性变更信息
*/
type TaobaoItemCatpropsModificationGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoItemCatpropsModificationGetResponse `json:"item_catprops_modification_get_response,omitempty"` 
    TaobaoItemCatpropsModificationGetResponse
}

/* model for simplify = false
type TaobaoItemCatpropsModificationGetResponse struct {

    // 返回结果
    
    Results  struct {
        PropsModificationResult  []PropsModificationResult `json:"props_modification_result,omitempty"`
    } `json:"results,omitempty"`
    

}
*/

type TaobaoItemCatpropsModificationGetResponse struct {

    // 返回结果
    Results   []PropsModificationResult `json:"results,omitempty"`

}
