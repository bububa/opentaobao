package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
调拨单创建 APIResponse
taobao.qimen.transferorder.create

调拨单创建
*/
type TaobaoQimenTransferorderCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_transferorder_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *TaobaoQimenTransferorderCreateStruct `json:"response,omitempty" xml:"