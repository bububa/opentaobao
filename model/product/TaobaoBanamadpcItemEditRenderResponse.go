package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑商品发布页 APIResponse
taobao.banamadpc.item.edit.render

巴拿马供应商通过此接口获取编辑商品发布页
*/
type TaobaoBanamadpcItemEditRenderAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"banamadpc_item_edit_render_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 无
    
    ApiResult   *TaobaoBanamadpcItemEditRenderApiResult `json:"api_result,omitempty" xml:"