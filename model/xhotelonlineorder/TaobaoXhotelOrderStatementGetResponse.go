package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询账单信息 APIResponse
taobao.xhotel.order.statement.get

阿里根据此接口定义输出订单账务明细，结账状态发生变化时阿里需推送账单信息。系统商可实时调用该接口来查询订单的详情
*/
type TaobaoXhotelOrderStatementGetAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelOrderStatementGetResponse
}

type TaobaoXhotelOrderStatementGetResponse struct {
    XMLName xml.Name `xml:"xhotel_order_statement_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 系统自动生成
    
    Results   []StatementOrder `json:"results,omitempty" xml:"results>statement_order,omitempty"`
    
    
}
