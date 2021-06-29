package degoperation

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通用用户抽奖次数限制 API返回值 
taobao.degoperation.get.by.eventkey

通用用户抽奖次数限制
*/
type TaobaoDegoperationGetByEventkeyAPIResponse struct {
    model.CommonResponse
    TaobaoDegoperationGetByEventkeyResponse
}

// 通用用户抽奖次数限制 成功返回结果
type TaobaoDegoperationGetByEventkeyResponse struct {
    XMLName xml.Name `xml:"degoperation_get_by_eventkey_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *BonusResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
