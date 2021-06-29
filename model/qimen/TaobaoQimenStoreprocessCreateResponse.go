package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
仓内加工单创建接口 APIResponse
taobao.qimen.storeprocess.create

ERP调用奇门的接口,创建仓内加工单
*/
type TaobaoQimenStoreprocessCreateAPIResponse struct {
    model.CommonResponse
    TaobaoQimenStoreprocessCreateResponse
}

type TaobaoQimenStoreprocessCreateResponse struct {
    XMLName xml.Name `xml:"qimen_storeprocess_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 
    
    Response   *StoreProcessCreateResponse `json:"response,omitempty" xml:"response,omitempty"`

    
}
