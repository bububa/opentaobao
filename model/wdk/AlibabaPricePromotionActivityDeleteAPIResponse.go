package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabapricepromotionactivitydeleteAPIResponse 删除档期活动 API返回值
// alibaba.price.promotion.activity.delete
//
// 删除盒马帮档期活动
type AlibabapricepromotionactivitydeleteAPIResponse struct {
	model.CommonResponse
	AlibabapricepromotionactivitydeleteAPIResponseModel
}

// AlibabapricepromotionactivitydeleteAPIResponseModel is 删除档期活动 成功返回结果
type AlibabapricepromotionactivitydeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_price_promotion_activity_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabapricepromotionactivitydeleteResult `json:"result,omitempty" xml:"result,omitempty"`
}
