package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家收到的退款列表 APIResponse
taobao.refunds.receive.get

查询卖家收到的退款列表
*/
type TaobaoRefundsReceiveGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoRefundsReceiveGetResponse `json:"taobao_refunds_receive_get_response,omitempty"`
}

type TaobaoRefundsReceiveGetResponse struct {

    // 搜索到的退款信息总数
    TotalResults   int64 `json:"total_results,omitempty"`

    // 是否存在下一页
    HasNext   bool `json:"has_next,omitempty"`

    // 搜索到的退款信息列表
    Refunds   []Refund `json:"refunds,omitempty"`

}
