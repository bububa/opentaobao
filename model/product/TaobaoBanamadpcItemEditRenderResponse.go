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
    Response *TaobaoBanamadpcItemEditRenderResponse `json:"taobao_banamadpc_item_edit_render_response,omitempty"`
}

type TaobaoBanamadpcItemEditRenderResponse struct {

    // 无
    ApiResult   *TaobaoBanamadpcItemEditRenderApiResult `json:"api_result,omitempty"`

}
