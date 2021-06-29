package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
信息流新增并且绑定创意 API返回值 
taobao.feedflow.item.adgroup.creative.add.bind

信息流新增并且绑定创意
*/
type TaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemAdgroupCreativeAddBindResponse
}

// 信息流新增并且绑定创意 成功返回结果
type TaobaoFeedflowItemAdgroupCreativeAddBindResponse struct {
    XMLName xml.Name `xml:"feedflow_item_adgroup_creative_add_bind_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果对象
    Result   *TaobaoFeedflowItemAdgroupCreativeAddBindResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
