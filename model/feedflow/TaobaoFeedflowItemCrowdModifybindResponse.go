package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
修改人群出价或状态 API返回值 
taobao.feedflow.item.crowd.modifybind

修改人群出价或状态
*/
type TaobaoFeedflowItemCrowdModifybindAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemCrowdModifybindResponse
}

// 修改人群出价或状态 成功返回结果
type TaobaoFeedflowItemCrowdModifybindResponse struct {
    XMLName xml.Name `xml:"feedflow_item_crowd_modifybind_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果对象
    Result   *TaobaoFeedflowItemCrowdModifybindResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
