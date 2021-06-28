package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新发商品发布页 APIResponse
taobao.banamadpc.item.render

巴拿马供应商通过此接口新发商品发布页
*/
type TaobaoBanamadpcItemRenderAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBanamadpcItemRenderResponse `json:"banamadpc_item_render_response,omitempty"` 
    TaobaoBanamadpcItemRenderResponse
}

/* model for simplify = false
type TaobaoBanamadpcItemRenderResponse struct {

    // 根据站点名称查询产品
    
    ApiResult  *struct {
        TaobaoBanamadpcItemRenderApiResult  *TaobaoBanamadpcItemRenderApiResult `json:"taobao_banamadpc_item_render_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type TaobaoBanamadpcItemRenderResponse struct {

    // 根据站点名称查询产品
    ApiResult   *TaobaoBanamadpcItemRenderApiResult `json:"api_result,omitempty"`

}
