package singletreasure

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosingletreasureactivityitembatchupdateAPIResponse 批量修改商品接口 API返回值
// taobao.singletreasure.activity.item.batchupdate
//
// 批量修改商品优惠接口
type TaobaosingletreasureactivityitembatchupdateAPIResponse struct {
	model.CommonResponse
	TaobaosingletreasureactivityitembatchupdateAPIResponseModel
}

// TaobaosingletreasureactivityitembatchupdateAPIResponseModel is 批量修改商品接口 成功返回结果
type TaobaosingletreasureactivityitembatchupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"singletreasure_activity_item_batchupdate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaosingletreasureactivityitembatchupdateResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
