package nlife

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取采购退货单列表 APIResponse
alibaba.nlife.b2b.trade.refund.list

获取采购退货单列表
*/
type AlibabaNlifeB2bTradeRefundListAPIResponse struct {
    model.CommonResponse
    AlibabaNlifeB2bTradeRefundListResponse
}

type AlibabaNlifeB2bTradeRefundListResponse struct {
    XMLName xml.Name `xml:"alibaba_nlife_b2b_trade_refund_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询结果数据
    
    Data   *RefundListResponseDo `json:"data,omitempty" xml:"data,omitempty"`

    
}
