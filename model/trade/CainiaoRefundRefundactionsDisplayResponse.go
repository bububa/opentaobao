package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
退货退款操作的展示信息(展现给买家) APIResponse
cainiao.refund.refundactions.display

退货退款操作的展示信息(展现给买家)
*/
type CainiaoRefundRefundactionsDisplayAPIResponse struct {
    model.CommonResponse
    CainiaoRefundRefundactionsDisplayResponse
}

type CainiaoRefundRefundactionsDisplayResponse struct {
    XMLName xml.Name `xml:"cainiao_refund_refundactions_display_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果对象
    
    Result   *CainiaoRefundRefundactionsDisplayBizResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
