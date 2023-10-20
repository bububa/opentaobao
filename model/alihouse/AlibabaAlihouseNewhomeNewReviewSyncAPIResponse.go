package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeNewReviewSyncAPIResponse 新测评基础信息接口 API返回值
// alibaba.alihouse.newhome.new.review.sync
//
// 新测评基础信息接口
type AlibabaAlihouseNewhomeNewReviewSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeNewReviewSyncAPIResponseModel
}

// AlibabaAlihouseNewhomeNewReviewSyncAPIResponseModel is 新测评基础信息接口 成功返回结果
type AlibabaAlihouseNewhomeNewReviewSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_new_review_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeNewReviewSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
