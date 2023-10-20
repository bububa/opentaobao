package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse 同步品活动标 API返回值
// alibaba.alihouse.newhome.tradeitem.features.sync
//
// 同步品活动标
type AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponseModel is 同步品活动标 成功返回结果
type AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_tradeitem_features_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeTradeitemFeaturesSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse
func GetAlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse() *AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse {
	return poolAlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse.Get().(*AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse 将 AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse(v *AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse.Put(v)
}
