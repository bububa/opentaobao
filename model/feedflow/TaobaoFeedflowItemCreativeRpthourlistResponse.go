package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
超级推荐【商品推广】创意分时报表查询 APIResponse
taobao.feedflow.item.creative.rpthourlist

创意分时数据查询，支持广告主查询最近90天内某一天的创意维度分时报表数据
*/
type TaobaoFeedflowItemCreativeRpthourlistAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemCreativeRpthourlistResponse
}

type TaobaoFeedflowItemCreativeRpthourlistResponse struct {
    XMLName xml.Name `xml:"feedflow_item_creative_rpthourlist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 异步获取历史数据接口返回结果
    
    Result   *TaobaoFeedflowItemCreativeRpthourlistResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
