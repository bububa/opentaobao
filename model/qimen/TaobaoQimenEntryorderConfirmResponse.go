package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
入库单确认接口 APIResponse
taobao.qimen.entryorder.confirm

WMS调用接口，回传入库单信息;
*/
type TaobaoQimenEntryorderConfirmAPIResponse struct {
    model.CommonResponse
    TaobaoQimenEntryorderConfirmResponse
}

type TaobaoQimenEntryorderConfirmResponse struct {
    XMLName xml.Name `xml:"qimen_entryorder_confirm_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 
    
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`

    
}
