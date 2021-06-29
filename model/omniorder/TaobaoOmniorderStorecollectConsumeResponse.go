package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道门店自提核销订单 APIResponse
taobao.omniorder.storecollect.consume

全渠道门店自提核销订单
*/
type TaobaoOmniorderStorecollectConsumeAPIResponse struct {
    model.CommonResponse
    TaobaoOmniorderStorecollectConsumeResponse
}

type TaobaoOmniorderStorecollectConsumeResponse struct {
    XMLName xml.Name `xml:"omniorder_storecollect_consume_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 0表示成功
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`

    
    // 核销错误信息
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`

    
}
