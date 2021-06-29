package feedflow

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取有权限的定向列表 API返回值 
taobao.feedflow.item.target.validlist

获取有权限的定向列表
*/
type TaobaoFeedflowItemTargetValidlistAPIResponse struct {
    model.CommonResponse
    TaobaoFeedflowItemTargetValidlistResponse
}

// 获取有权限的定向列表 成功返回结果
type TaobaoFeedflowItemTargetValidlistResponse struct {
    XMLName xml.Name `xml:"feedflow_item_target_validlist_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果对象
    Result   *TaobaoFeedflowItemTargetValidlistResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
