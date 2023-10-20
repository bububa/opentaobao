package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeNewRichReviewSyncAPIResponse 新评测改造接口同步 API返回值
// alibaba.alihouse.newhome.new.rich.review.sync
//
// 新评测改造接口同步
type AlibabaAlihouseNewhomeNewRichReviewSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeNewRichReviewSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeNewRichReviewSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeNewRichReviewSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeNewRichReviewSyncAPIResponseModel is 新评测改造接口同步 成功返回结果
type AlibabaAlihouseNewhomeNewRichReviewSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_new_rich_review_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果实体类
	Result *AlibabaAlihouseNewhomeNewRichReviewSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeNewRichReviewSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeNewRichReviewSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeNewRichReviewSyncAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeNewRichReviewSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeNewRichReviewSyncAPIResponse
func GetAlibabaAlihouseNewhomeNewRichReviewSyncAPIResponse() *AlibabaAlihouseNewhomeNewRichReviewSyncAPIResponse {
	return poolAlibabaAlihouseNewhomeNewRichReviewSyncAPIResponse.Get().(*AlibabaAlihouseNewhomeNewRichReviewSyncAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeNewRichReviewSyncAPIResponse 将 AlibabaAlihouseNewhomeNewRichReviewSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeNewRichReviewSyncAPIResponse(v *AlibabaAlihouseNewhomeNewRichReviewSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeNewRichReviewSyncAPIResponse.Put(v)
}
