package singletreasure

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosingletreasureactivityitembatchaddAPIResponse 批量添加商品接口 API返回值
// taobao.singletreasure.activity.item.batchadd
//
// 向活动中批量添加商品优惠
type TaobaosingletreasureactivityitembatchaddAPIResponse struct {
	model.CommonResponse
	TaobaosingletreasureactivityitembatchaddAPIResponseModel
}

// TaobaosingletreasureactivityitembatchaddAPIResponseModel is 批量添加商品接口 成功返回结果
type TaobaosingletreasureactivityitembatchaddAPIResponseModel struct {
	XMLName xml.Name `xml:"singletreasure_activity_item_batchadd_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaosingletreasureactivityitembatchaddResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
