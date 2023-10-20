package singletreasure

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosingletreasureactivitynamequeryAPIResponse 查询官方的活动名称接口 API返回值
// taobao.singletreasure.activity.name.query
//
// 查询官方的活动名称列表接口
type TaobaosingletreasureactivitynamequeryAPIResponse struct {
	model.CommonResponse
	TaobaosingletreasureactivitynamequeryAPIResponseModel
}

// TaobaosingletreasureactivitynamequeryAPIResponseModel is 查询官方的活动名称接口 成功返回结果
type TaobaosingletreasureactivitynamequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"singletreasure_activity_name_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaosingletreasureactivitynamequeryResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
