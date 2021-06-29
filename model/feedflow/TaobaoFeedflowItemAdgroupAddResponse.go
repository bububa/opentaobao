package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流增加单元 API返回值 
taobao.feedflow.item.adgroup.add

信息流增加单元
*/
type TaobaoFeedflowItemAdgroupAddAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemAdgroupAddResponse
}

// 信息流增加单元 成功返回结果
type TaobaoFeedflowItemAdgroupAddResponse struct {
    XMLName xml.Name `xml:"feedflow_item_adgroup_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *TaobaoFeedflowItemAdgroupAddResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
