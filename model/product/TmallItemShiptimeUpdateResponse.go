package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新商品发货时间 APIResponse
tmall.item.shiptime.update

增加更新删除商品/SKU发货时间(支持同一商品下的SKU同时批量更新)
1.
    {
        "shipTimeType": 2,  ----相对发货时间（值为1则为绝对发货时间）
        "updateType": 0 ---更新SKU
    },

    按照指定SKU更新指定SKU的发货时间，如果原本是商品级发货时间，商品级发货时间也清空

2.
    {
        "shipTimeType": 0, -- 删除发货时间
        "updateType": 0 --更新SKU
    },
    按照指定SKU删除指定SKU的发货时间

3.

    {
        "shipTimeType": 2,  ----相对发货时间（值为1则为绝对发货时间）
        "updateType": 1 ---更新商品
    },

    更新商品级发货时间，如果原本是SKU级发货时间，清空所有SKU上的发货时间

4.
    {
        "shipTimeType": 0, -- 删除发货时间
        "updateType": 1 --更新商品
    },
    删除商品级的发货时间
*/
type TmallItemShiptimeUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TmallItemShiptimeUpdateResponse `json:"tmall_item_shiptime_update_response,omitempty"` 
    TmallItemShiptimeUpdateResponse
}

/* model for simplify = false
type TmallItemShiptimeUpdateResponse struct {

    // 被修改的商品ID
    
    Result   string `json:"result,omitempty"`
    

}
*/

type TmallItemShiptimeUpdateResponse struct {

    // 被修改的商品ID
    Result   string `json:"result,omitempty"`

}
