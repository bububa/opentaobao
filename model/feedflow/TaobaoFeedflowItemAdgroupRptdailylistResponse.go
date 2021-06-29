package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推广单元分日数据查询 APIResponse
taobao.feedflow.item.adgroup.rptdailylist

推广单元分日数据查询
*/
type TaobaoFeedflowItemAdgroupRptdailylistAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemAdgroupRptdailylistResponse
}

type TaobaoFeedflowItemAdgroupRptdailylistResponse struct {
    XMLName xml.Name `xml:"feedflow_item_adgroup_rptdailylist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TaobaoFeedflowItemAdgroupRptdailylistResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
