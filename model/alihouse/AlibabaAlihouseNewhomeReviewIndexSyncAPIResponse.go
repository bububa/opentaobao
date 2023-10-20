package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeReviewIndexSyncAPIResponse 新测评乐居指数接口 API返回值
// alibaba.alihouse.newhome.review.index.sync
//
// 新测评乐居指数同步数据
type AlibabaAlihouseNewhomeReviewIndexSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeReviewIndexSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeReviewIndexSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeReviewIndexSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeReviewIndexSyncAPIResponseModel is 新测评乐居指数接口 成功返回结果
type AlibabaAlihouseNewhomeReviewIndexSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_review_index_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeReviewIndexSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeReviewIndexSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeReviewIndexSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeReviewIndexSyncAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeReviewIndexSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeReviewIndexSyncAPIResponse
func GetAlibabaAlihouseNewhomeReviewIndexSyncAPIResponse() *AlibabaAlihouseNewhomeReviewIndexSyncAPIResponse {
	return poolAlibabaAlihouseNewhomeReviewIndexSyncAPIResponse.Get().(*AlibabaAlihouseNewhomeReviewIndexSyncAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeReviewIndexSyncAPIResponse 将 AlibabaAlihouseNewhomeReviewIndexSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeReviewIndexSyncAPIResponse(v *AlibabaAlihouseNewhomeReviewIndexSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeReviewIndexSyncAPIResponse.Put(v)
}
