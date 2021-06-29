package openmall

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取openmall订单 APIResponse
taobao.openmall.trade.batch.get

批量获取openmall订单
注意：该接口数据存在延迟，实时数据请通过taobao.openmall.trade.get获取
*/
type TaobaoOpenmallTradeBatchGetAPIResponse struct {
    model.CommonResponse
    TaobaoOpenmallTradeBatchGetResponse
}

type TaobaoOpenmallTradeBatchGetResponse struct {
    XMLName xml.Name `xml:"openmall_trade_batch_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 订单列表
    
    Entities   []TopTradeDetailVo `json:"entities,omitempty" xml:"entities>top_trade_detail_vo,omitempty"`
    
    
    // 范围内总订单数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`

    
}
