package openmall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取openmall退款单 API返回值 
taobao.openmall.refund.batch.get

批量获取openmall退款单
注意：该接口信息存在延迟，如需实时详情请访问taobao.openmall.refund.get
*/
type TaobaoOpenmallRefundBatchGetAPIResponse struct {
    model.CommonResponse
    TaobaoOpenmallRefundBatchGetAPIResponseModel
}

// 批量获取openmall退款单 成功返回结果
type TaobaoOpenmallRefundBatchGetAPIResponseModel struct {
    XMLName xml.Name `xml:"openmall_refund_batch_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 退款列表
    Entities   []TopRefundVo `json:"entities,omitempty" xml:"entities>top_refund_vo,omitempty"`
    // 范围内总的退款单个数
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
