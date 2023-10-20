package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponse 组合货品新建&更新 API返回值
// alibaba.dchain.aoxiang.combineitem.batch.update.async
//
// 组合货品新建&amp;更新
type AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponseModel is 组合货品新建&更新 成功返回结果
type AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_combineitem_batch_update_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求出参
	CombineItemUpsertResponse *CombineItemUpsertAsyncResponse `json:"combine_item_upsert_response,omitempty" xml:"combine_item_upsert_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CombineItemUpsertResponse = nil
}

var poolAlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponse)
	},
}

// GetAlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponse
func GetAlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponse() *AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponse {
	return poolAlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponse.Get().(*AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponse)
}

// ReleaseAlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponse 将 AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponse(v *AlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangCombineitemBatchUpdateAsyncAPIResponse.Put(v)
}
