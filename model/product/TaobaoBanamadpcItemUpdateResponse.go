package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑商品 APIResponse
taobao.banamadpc.item.update

巴拿马供应商通过此接口编辑商品
*/
type TaobaoBanamadpcItemUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoBanamadpcItemUpdateResponse
}

type TaobaoBanamadpcItemUpdateResponse struct {
    XMLName xml.Name `xml:"banamadpc_item_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 无
    
    ApiResult   *TaobaoBanamadpcItemUpdateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`

    
}
