package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据单元id删除单元 API返回值 
taobao.feedflow.item.adgroup.delete

根据单元id删除单元
*/
type TaobaoFeedflowItemAdgroupDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemAdgroupDeleteResponse
}

// 根据单元id删除单元 成功返回结果
type TaobaoFeedflowItemAdgroupDeleteResponse struct {
    XMLName xml.Name `xml:"feedflow_item_adgroup_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *TaobaoFeedflowItemAdgroupDeleteResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
