package nlife

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取采购退货单列表 API返回值 
alibaba.nlife.b2b.trade.refund.list

获取采购退货单列表
*/
type AlibabaNlifeB2bTradeRefundListAPIResponse struct {
    model.CommonResponse
    AlibabaNlifeB2bTradeRefundListResponse
}

// 获取采购退货单列表 成功返回结果
type AlibabaNlifeB2bTradeRefundListResponse struct {
    XMLName xml.Name `xml:"alibaba_nlife_b2b_trade_refund_list_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 查询结果数据
    Data   *RefundListResponseDo `json:"data,omitempty" xml:"data,omitempty"`
}
