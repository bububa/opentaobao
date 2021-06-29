package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑商品发布页 API返回值 
taobao.banamadpc.item.edit.render

巴拿马供应商通过此接口获取编辑商品发布页
*/
type TaobaoBanamadpcItemEditRenderAPIResponse struct {
    model.CommonResponse
    TaobaoBanamadpcItemEditRenderResponse
}

// 编辑商品发布页 成功返回结果
type TaobaoBanamadpcItemEditRenderResponse struct {
    XMLName xml.Name `xml:"banamadpc_item_edit_render_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 无
    ApiResult   *TaobaoBanamadpcItemEditRenderApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
