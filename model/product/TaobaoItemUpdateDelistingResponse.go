package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商品下架 APIResponse
taobao.item.update.delisting

* 单个商品下架<br/>    * 输入的num_iid必须属于当前会话用户
*/
type TaobaoItemUpdateDelistingAPIResponse struct {
    model.CommonResponse
    Response *TaobaoItemUpdateDelistingResponse `json:"taobao_item_update_delisting_response,omitempty"`
}

type TaobaoItemUpdateDelistingResponse struct {

    // 返回商品更新信息：返回的结果是:num_iid和modified
    Item   *Item `json:"item,omitempty"`

}
