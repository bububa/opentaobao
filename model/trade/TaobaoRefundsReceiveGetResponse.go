package trade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卖家收到的退款列表 APIResponse
taobao.refunds.receive.get

查询卖家收到的退款列表
*/
type TaobaoRefundsReceiveGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"refunds_receive_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 搜索到的退款信息总数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"