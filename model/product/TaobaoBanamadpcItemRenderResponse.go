package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新发商品发布页 APIResponse
taobao.banamadpc.item.render

巴拿马供应商通过此接口新发商品发布页
*/
type TaobaoBanamadpcItemRenderAPIResponse struct {
    model.CommonResponse
    TaobaoBanamadpcItemRenderResponse
}

type TaobaoBanamadpcItemRenderResponse struct {
    XMLName xml.Name `xml:"banamadpc_item_render_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 根据站点名称查询产品
    
    ApiResult   *TaobaoBanamadpcItemRenderApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`

    
}
