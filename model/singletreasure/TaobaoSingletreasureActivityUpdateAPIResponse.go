package singletreasure

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosingletreasureactivityupdateAPIResponse 修改活动接口 API返回值
// taobao.singletreasure.activity.update
//
// 修改活动接口
type TaobaosingletreasureactivityupdateAPIResponse struct {
	model.CommonResponse
	TaobaosingletreasureactivityupdateAPIResponseModel
}

// TaobaosingletreasureactivityupdateAPIResponseModel is 修改活动接口 成功返回结果
type TaobaosingletreasureactivityupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"singletreasure_activity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaosingletreasureactivityupdateResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
