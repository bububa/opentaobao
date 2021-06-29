package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
判断该订单能执行的逆向操作集合列表 APIResponse
cainiao.refund.refundactions.get

判断该订单能执行的逆向操作集合列表
*/
type CainiaoRefundRefundactionsGetAPIResponse struct {
    model.CommonResponse
    CainiaoRefundRefundactionsGetResponse
}

type CainiaoRefundRefundactionsGetResponse struct {
    XMLName xml.Name `xml:"cainiao_refund_refundactions_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果对象
    
    Result   *CainiaoRefundRefundactionsGetBizResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
