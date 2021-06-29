package eticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证操作日志查询 APIResponse
taobao.vmarket.eticket.oplogs.get

电子凭证核销日志查询
*/
type TaobaoVmarketEticketOplogsGetAPIResponse struct {
    model.CommonResponse
    TaobaoVmarketEticketOplogsGetResponse
}

type TaobaoVmarketEticketOplogsGetResponse struct {
    XMLName xml.Name `xml:"vmarket_eticket_oplogs_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 符合条件的记录总数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`

    
    // 操作日志列表
    
    EticketOpLogs   []EticketOpLog `json:"eticket_op_logs,omitempty" xml:"eticket_op_logs>eticket_op_log,omitempty"`
    
    
}
