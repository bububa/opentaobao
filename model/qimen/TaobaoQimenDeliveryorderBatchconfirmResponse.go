package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单确认接口 APIResponse
taobao.qimen.deliveryorder.batchconfirm

taobao.qimen.deliveryorder.batchconfirm
*/
type TaobaoQimenDeliveryorderBatchconfirmAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_deliveryorder_batchconfirm_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *Response `json:"response,omitempty" xml:"