package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
仓内加工单确认接口 APIResponse
taobao.qimen.storeprocess.confirm

WMS调用奇门的接口,回传仓内加工单创建情况
*/
type TaobaoQimenStoreprocessConfirmAPIResponse struct {
    model.CommonResponse
    TaobaoQimenStoreprocessConfirmResponse
}

type TaobaoQimenStoreprocessConfirmResponse struct {
    XMLName xml.Name `xml:"qimen_storeprocess_confirm_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 
    
    Response   *Response `json:"response,omitempty" xml:"response,omitempty"`

    
}
