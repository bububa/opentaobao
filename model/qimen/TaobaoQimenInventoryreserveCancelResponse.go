package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
库存预占取消接口 APIResponse
taobao.qimen.inventoryreserve.cancel

库存预占取消
*/
type TaobaoQimenInventoryreserveCancelAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_inventoryreserve_cancel_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *Response `json:"response,omitempty" xml:"