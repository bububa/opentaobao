package elife

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌惠卡券冲正退还 APIResponse
taobao.elife.lifecard.refund

淘宝生活汇消费卡虚拟卡，线下冲正退货接口
*/
type TaobaoElifeLifecardRefundAPIResponse struct {
    model.CommonResponse
    TaobaoElifeLifecardRefundResponse
}

type TaobaoElifeLifecardRefundResponse struct {
    XMLName xml.Name `xml:"elife_lifecard_refund_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回码，成功为空
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 本金
    
    Amount   int64 `json:"amount,omitempty" xml:"amount,omitempty"`

    
    // 膨胀金
    
    InflateAmount   int64 `json:"inflate_amount,omitempty" xml:"inflate_amount,omitempty"`

    
    // 返回信息
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // 成功失败标志
    
    Successed   bool `json:"successed,omitempty" xml:"successed,omitempty"`

    
}
