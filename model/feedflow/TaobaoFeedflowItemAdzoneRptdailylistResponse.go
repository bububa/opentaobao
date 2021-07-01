package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
资源包分日数据查询 API返回值 
taobao.feedflow.item.adzone.rptdailylist

资源包分日数据查询
*/
type TaobaoFeedflowItemAdzoneRptdailylistAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemAdzoneRptdailylistResponse
}

// 资源包分日数据查询 成功返回结果
type TaobaoFeedflowItemAdzoneRptdailylistResponse struct {
    XMLName xml.Name `xml:"feedflow_item_adzone_rptdailylist_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *TaobaoFeedflowItemAdzoneRptdailylistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
