package paimai

import (
    "github.com/bububa/opentaobao/model"
)

/* 
拍卖相关类目属性 APIResponse
taobao.paimai.itemprops.get

读取拍卖相关类目属性
*/
type TaobaoPaimaiItempropsGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoPaimaiItempropsGetResponse `json:"paimai_itemprops_get_response,omitempty"` 
    TaobaoPaimaiItempropsGetResponse
}

/* model for simplify = false
type TaobaoPaimaiItempropsGetResponse struct {

    // 类目属性信息(如果是取全量或者增量，不包括属性值),根据fields传入的参数返回相应的结果
    
    ItemProps  struct {
        ItemProp  []ItemProp `json:"item_prop,omitempty"`
    } `json:"item_props,omitempty"`
    

    // lastModified
    
    LastModified   string `json:"last_modified,omitempty"`
    

}
*/

type TaobaoPaimaiItempropsGetResponse struct {

    // 类目属性信息(如果是取全量或者增量，不包括属性值),根据fields传入的参数返回相应的结果
    ItemProps   []ItemProp `json:"item_props,omitempty"`

    // lastModified
    LastModified   string `json:"last_modified,omitempty"`

}
