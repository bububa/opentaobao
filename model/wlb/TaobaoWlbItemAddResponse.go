package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
添加单个物流宝商品 APIResponse
taobao.wlb.item.add

添加物流宝商品，支持物流宝子商品和属性添加
*/
type TaobaoWlbItemAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoWlbItemAddResponse `json:"taobao_wlb_item_add_response,omitempty"`
}

type TaobaoWlbItemAddResponse struct {

    // 新增的商品
    ItemId   string `json:"item_id,omitempty"`

}
