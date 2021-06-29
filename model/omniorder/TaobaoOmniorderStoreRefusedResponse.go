package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
Pos端门店拒单 APIResponse
taobao.omniorder.store.refused

ISV Pos端门店拒单，通知星盘
*/
type TaobaoOmniorderStoreRefusedAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderStoreRefusedResponse
}

type TaobaoOmniorderStoreRefusedResponse struct {
    XMLName xml.Name `xml:"omniorder_store_refused_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 正常为0,其他表示异常
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`

    
    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
}
