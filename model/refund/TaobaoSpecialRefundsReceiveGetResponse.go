package refund

import (
    "github.com/bububa/opentaobao/model"
)

/* 
特殊退款类型的纠纷单列表查询 APIResponse
taobao.special.refunds.receive.get

特殊退款类型的纠纷单列表查询
*/
type TaobaoSpecialRefundsReceiveGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSpecialRefundsReceiveGetResponse `json:"taobao_special_refunds_receive_get_response,omitempty"`
}

type TaobaoSpecialRefundsReceiveGetResponse struct {

    // 搜索到的退款信息总数
    TotalResults   int64 `json:"total_results,omitempty"`

    // 是否存在下一页
    HasNext   bool `json:"has_next,omitempty"`

    // 搜索到的退款信息列表
    Refunds   []Refund `json:"refunds,omitempty"`

}
