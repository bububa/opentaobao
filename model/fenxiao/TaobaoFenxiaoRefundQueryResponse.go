package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量查询采购退款 APIResponse
taobao.fenxiao.refund.query

供应商按查询条件批量查询代销采购退款
*/
type TaobaoFenxiaoRefundQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"fenxiao_refund_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 按查询条件查到的记录总数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"