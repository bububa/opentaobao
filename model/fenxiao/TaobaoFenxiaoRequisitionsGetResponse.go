package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
合作申请查询 APIResponse
taobao.fenxiao.requisitions.get

合作申请查询
*/
type TaobaoFenxiaoRequisitionsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoFenxiaoRequisitionsGetResponse `json:"taobao_fenxiao_requisitions_get_response,omitempty"`
}

type TaobaoFenxiaoRequisitionsGetResponse struct {

    // 操作是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 结果记录数
    TotalResults   int64 `json:"total_results,omitempty"`

    // 合作申请
    Requisitions   []Requisition `json:"requisitions,omitempty"`

}
