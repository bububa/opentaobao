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
    TaobaoRefundsReceiveGetResponse
}

type TaobaoRefundsReceiveGetResponse struct {
    XMLName xml.Name `xml:"refunds_receive_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 搜索到的退款信息总数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`

    
    // 是否存在下一页
    
    HasNext   bool `json:"has_next,omitempty" xml:"has_next,omitempty"`

    
    // 搜索到的退款信息列表
    
    Refunds   []Refund `json:"refunds,omitempty" xml:"refunds>refund,omitempty"`
    
    
}
