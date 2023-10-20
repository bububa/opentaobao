package singletreasure

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosingletreasureactivityitemupdateAPIResponse 更新单品优惠接口 API返回值
// taobao.singletreasure.activity.item.update
//
// 更新单品优惠接口
type TaobaosingletreasureactivityitemupdateAPIResponse struct {
	model.CommonResponse
	TaobaosingletreasureactivityitemupdateAPIResponseModel
}

// TaobaosingletreasureactivityitemupdateAPIResponseModel is 更新单品优惠接口 成功返回结果
type TaobaosingletreasureactivityitemupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"singletreasure_activity_item_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaosingletreasureactivityitemupdateResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
