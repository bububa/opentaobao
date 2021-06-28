package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
调拨单查询 APIResponse
taobao.qimen.transferorder.query

调拨单查询
*/
type TaobaoQimenTransferorderQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qimen_transferorder_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *TaobaoQimenTransferorderQueryStruct `json:"response,omitempty" xml:"