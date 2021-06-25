package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
taobao.item.update.delisting.tmall APIResponse
taobao.item.update.delisting.tmall

* 单个商品下架<br/>    * 输入的num_iid必须属于当前会话用户
*/
type TaobaoItemUpdateDelistingTmallAPIResponse struct {
    model.CommonResponse
    Response *TaobaoItemUpdateDelistingTmallResponse `json:"taobao_item_update_delisting_tmall_response,omitempty"`
}

type TaobaoItemUpdateDelistingTmallResponse struct {

    // 返回商品更新信息：返回的结果是:num_iid和modified
    Item   *Item `json:"item,omitempty"`

}
