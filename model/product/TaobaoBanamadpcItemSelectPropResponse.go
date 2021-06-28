package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取子属性 APIResponse
taobao.banamadpc.item.select.prop

巴拿马供应商通过此接口获取子属性
*/
type TaobaoBanamadpcItemSelectPropAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBanamadpcItemSelectPropResponse `json:"banamadpc_item_select_prop_response,omitempty"` 
    TaobaoBanamadpcItemSelectPropResponse
}

/* model for simplify = false
type TaobaoBanamadpcItemSelectPropResponse struct {

    // 无
    
    ApiResult  *struct {
        TaobaoBanamadpcItemSelectPropApiResult  *TaobaoBanamadpcItemSelectPropApiResult `json:"taobao_banamadpc_item_select_prop_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type TaobaoBanamadpcItemSelectPropResponse struct {

    // 无
    ApiResult   *TaobaoBanamadpcItemSelectPropApiResult `json:"api_result,omitempty"`

}
