package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创意分日数据查询 APIResponse
taobao.feedflow.item.creative.rptdailylist

创意分日数据查询
*/
type TaobaoFeedflowItemCreativeRptdailylistAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemCreativeRptdailylistResponse
}

type TaobaoFeedflowItemCreativeRptdailylistResponse struct {
    XMLName xml.Name `xml:"feedflow_item_creative_rptdailylist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TaobaoFeedflowItemCreativeRptdailylistResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
