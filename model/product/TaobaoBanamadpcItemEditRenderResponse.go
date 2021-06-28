package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
编辑商品发布页 APIResponse
taobao.banamadpc.item.edit.render

巴拿马供应商通过此接口获取编辑商品发布页
*/
type TaobaoBanamadpcItemEditRenderAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBanamadpcItemEditRenderResponse `json:"banamadpc_item_edit_render_response,omitempty"` 
    TaobaoBanamadpcItemEditRenderResponse
}

/* model for simplify = false
type TaobaoBanamadpcItemEditRenderResponse struct {

    // 无
    
    ApiResult  *struct {
        TaobaoBanamadpcItemEditRenderApiResult  *TaobaoBanamadpcItemEditRenderApiResult `json:"taobao_banamadpc_item_edit_render_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type TaobaoBanamadpcItemEditRenderResponse struct {

    // 无
    ApiResult   *TaobaoBanamadpcItemEditRenderApiResult `json:"api_result,omitempty"`

}
