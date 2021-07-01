package degoperation

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
活动中奖记录 API返回值 
taobao.degoperation.show.top.records

活动中奖记录
*/
type TaobaoDegoperationShowTopRecordsAPIResponse struct {
    model.CommonResponse
    TaobaoDegoperationShowTopRecordsAPIResponseModel
}

// 活动中奖记录 成功返回结果
type TaobaoDegoperationShowTopRecordsAPIResponseModel struct {
    XMLName xml.Name `xml:"degoperation_show_top_records_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *BonusResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
