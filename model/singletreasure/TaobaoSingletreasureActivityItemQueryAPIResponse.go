package singletreasure

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosingletreasureactivityitemqueryAPIResponse 查询活动下的优惠信息 API返回值
// taobao.singletreasure.activity.item.query
//
// 分页查询活动下的商品优惠信息
type TaobaosingletreasureactivityitemqueryAPIResponse struct {
	model.CommonResponse
	TaobaosingletreasureactivityitemqueryAPIResponseModel
}

// TaobaosingletreasureactivityitemqueryAPIResponseModel is 查询活动下的优惠信息 成功返回结果
type TaobaosingletreasureactivityitemqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"singletreasure_activity_item_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *PageResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
