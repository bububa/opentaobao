package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退货入库单确认接口 APIResponse
taobao.qimen.returnorder.confirm

taobao.qimen.returnorder.confirm
*/
type TaobaoQimenReturnorderConfirmAPIResponse struct {
    model.CommonResponse
    TaobaoQimenReturnorderConfirmResponse
}

type TaobaoQimenReturnorderConfirmResponse struct {
    XMLName xml.Name `xml:"qimen_returnorder_confirm_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 
    
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`

    
}
