package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取子属性 APIResponse
taobao.banamadpc.item.select.prop

巴拿马供应商通过此接口获取子属性
*/
type TaobaoBanamadpcItemSelectPropAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"banamadpc_item_select_prop_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 无
    
    ApiResult   *TaobaoBanamadpcItemSelectPropApiResult `json:"api_result,omitempty" xml:"