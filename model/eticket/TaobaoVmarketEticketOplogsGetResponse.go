package eticket

import (
    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证操作日志查询 APIResponse
taobao.vmarket.eticket.oplogs.get

电子凭证核销日志查询
*/
type TaobaoVmarketEticketOplogsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoVmarketEticketOplogsGetResponse `json:"taobao_vmarket_eticket_oplogs_get_response,omitempty"`
}

type TaobaoVmarketEticketOplogsGetResponse struct {

    // 符合条件的记录总数
    TotalResults   int64 `json:"total_results,omitempty"`

    // 操作日志列表
    EticketOpLogs   []EticketOpLog `json:"eticket_op_logs,omitempty"`

}
