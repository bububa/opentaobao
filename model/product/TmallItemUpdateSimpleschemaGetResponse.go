package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
官网同购编辑商品的get接口 APIResponse
tmall.item.update.simpleschema.get

官网同购编辑商品的get接口
*/
type TmallItemUpdateSimpleschemaGetAPIResponse struct {
    model.CommonResponse
    Response *TmallItemUpdateSimpleschemaGetResponse `json:"tmall_item_update_simpleschema_get_response,omitempty"`
}

type TmallItemUpdateSimpleschemaGetResponse struct {

    // 返回结果
    Error   bool `json:"error,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 商品信息
    Result   string `json:"result,omitempty"`

}
