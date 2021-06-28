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
    TaobaoQimenTransferorderCreateResponse
}

type TaobaoQimenTransferorderCreateResponse struct {
    XMLName xml.Name `xml:"qimen_transferorder_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 
    
    Response   *TaobaoQimenTransferorderCreateStruct `json:"response,omitempty" xml:"response,omitempty"`

    
}
