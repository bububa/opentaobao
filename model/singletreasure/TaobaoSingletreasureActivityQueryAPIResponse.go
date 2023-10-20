package singletreasure

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosingletreasureactivityqueryAPIResponse 查询活动列表接口 API返回值
// taobao.singletreasure.activity.query
//
// 查询活动列表接口
type TaobaosingletreasureactivityqueryAPIResponse struct {
	model.CommonResponse
	TaobaosingletreasureactivityqueryAPIResponseModel
}

// TaobaosingletreasureactivityqueryAPIResponseModel is 查询活动列表接口 成功返回结果
type TaobaosingletreasureactivityqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"singletreasure_activity_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaosingletreasureactivityqueryResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
