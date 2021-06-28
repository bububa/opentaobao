package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新发商品 APIResponse
taobao.banamadpc.item.add

巴拿马供应商通过此接口新发商品
*/
type TaobaoBanamadpcItemAddAPIResponse struct {
    model.CommonResponse
    TaobaoBanamadpcItemAddResponse
}

type TaobaoBanamadpcItemAddResponse struct {
    XMLName xml.Name `xml:"banamadpc_item_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 无
    
    ApiResult   *TaobaoBanamadpcItemAddApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`

    
}
