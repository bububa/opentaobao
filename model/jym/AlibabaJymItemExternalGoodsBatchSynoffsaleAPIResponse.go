package jym

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponse 同步下架接口 API返回值
// alibaba.jym.item.external.goods.batch.synoffsale
//
// 同步下架接口
type AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponse struct {
	model.CommonResponse
	AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponseModel).Reset()
}

// AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponseModel is 同步下架接口 成功返回结果
type AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_item_external_goods_batch_synoffsale_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponse)
	},
}

// GetAlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponse 从 sync.Pool 获取 AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponse
func GetAlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponse() *AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponse {
	return poolAlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponse.Get().(*AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponse)
}

// ReleaseAlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponse 将 AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponse 保存到 sync.Pool
func ReleaseAlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponse(v *AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponse) {
	v.Reset()
	poolAlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponse.Put(v)
}
