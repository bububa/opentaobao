package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询单元列表 API返回值 
taobao.feedflow.item.adgroup.page

通过计划id查询单元信息
*/
type TaobaoFeedflowItemAdgroupPageAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemAdgroupPageResponse
}

// 查询单元列表 成功返回结果
type TaobaoFeedflowItemAdgroupPageResponse struct {
    XMLName xml.Name `xml:"feedflow_item_adgroup_page_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回数据
    Result   *ResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
