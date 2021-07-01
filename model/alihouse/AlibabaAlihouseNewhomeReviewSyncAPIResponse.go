package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihouseNewhomeReviewSyncAPIResponse
天猫好房楼盘评测同步 API返回值
alibaba.alihouse.newhome.review.sync

接受楼盘测评图文信息 */
type AlibabaAlihouseNewhomeReviewSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeReviewSyncAPIResponseModel
}

// AlibabaAlihouseNewhomeReviewSyncAPIResponseModel is 天猫好房楼盘评测同步 成功返回结果
type AlibabaAlihouseNewhomeReviewSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_review_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeReviewSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
