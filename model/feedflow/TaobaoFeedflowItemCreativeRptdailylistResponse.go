package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创意分日数据查询 API返回值 
taobao.feedflow.item.creative.rptdailylist

创意分日数据查询
*/
type TaobaoFeedflowItemCreativeRptdailylistAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemCreativeRptdailylistResponse
}

// 创意分日数据查询 成功返回结果
type TaobaoFeedflowItemCreativeRptdailylistResponse struct {
    XMLName xml.Name `xml:"feedflow_item_creative_rptdailylist_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *ResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
