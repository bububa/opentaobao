package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流单元内绑定资源位 API返回值 
taobao.feedflow.item.adgroup.adzone.bind

信息流单元内绑定资源位
*/
type TaobaoFeedflowItemAdgroupAdzoneBindAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemAdgroupAdzoneBindAPIResponseModel
}

// 信息流单元内绑定资源位 成功返回结果
type TaobaoFeedflowItemAdgroupAdzoneBindAPIResponseModel struct {
    XMLName xml.Name `xml:"feedflow_item_adgroup_adzone_bind_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果对象
    Result   *TaobaoFeedflowItemAdgroupAdzoneBindResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
