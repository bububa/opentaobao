package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新发商品 APIResponse
taobao.banamadpc.item.add

巴拿马供应商通过此接口新发商品
*/
type TaobaoBanamadpcItemAddAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBanamadpcItemAddResponse `json:"banamadpc_item_add_response,omitempty"` 
    TaobaoBanamadpcItemAddResponse
}

/* model for simplify = false
type TaobaoBanamadpcItemAddResponse struct {

    // 无
    
    ApiResult  *struct {
        TaobaoBanamadpcItemAddApiResult  *TaobaoBanamadpcItemAddApiResult `json:"taobao_banamadpc_item_add_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type TaobaoBanamadpcItemAddResponse struct {

    // 无
    ApiResult   *TaobaoBanamadpcItemAddApiResult `json:"api_result,omitempty"`

}
