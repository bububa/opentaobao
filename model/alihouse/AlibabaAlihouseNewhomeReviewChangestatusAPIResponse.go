package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeReviewChangestatusAPIResponse 楼盘测评草稿状态同步 API返回值
// alibaba.alihouse.newhome.review.changestatus
//
// 楼盘测评草稿状态更新
type AlibabaAlihouseNewhomeReviewChangestatusAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeReviewChangestatusAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeReviewChangestatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeReviewChangestatusAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeReviewChangestatusAPIResponseModel is 楼盘测评草稿状态同步 成功返回结果
type AlibabaAlihouseNewhomeReviewChangestatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_review_changestatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeReviewChangestatusResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeReviewChangestatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeReviewChangestatusAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeReviewChangestatusAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeReviewChangestatusAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeReviewChangestatusAPIResponse
func GetAlibabaAlihouseNewhomeReviewChangestatusAPIResponse() *AlibabaAlihouseNewhomeReviewChangestatusAPIResponse {
	return poolAlibabaAlihouseNewhomeReviewChangestatusAPIResponse.Get().(*AlibabaAlihouseNewhomeReviewChangestatusAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeReviewChangestatusAPIResponse 将 AlibabaAlihouseNewhomeReviewChangestatusAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeReviewChangestatusAPIResponse(v *AlibabaAlihouseNewhomeReviewChangestatusAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeReviewChangestatusAPIResponse.Put(v)
}
