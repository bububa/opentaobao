package alihouse

import (
	"encoding/xml"

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

// AlibabaAlihouseNewhomeReviewIndexSyncAPIResponseModel is 新测评乐居指数接口 成功返回结果
type AlibabaAlihouseNewhomeReviewIndexSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_review_index_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeReviewIndexSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
