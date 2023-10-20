package alihealthcrm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFmhealthButlerEnergySyncAPIResponse 同步用户消耗能量 API返回值
// alibaba.fmhealth.butler.energy.sync
//
// 同步用户消耗能量，用户消耗s点或卡路里后，同步给健康平台
type AlibabaFmhealthButlerEnergySyncAPIResponse struct {
	model.CommonResponse
	AlibabaFmhealthButlerEnergySyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaFmhealthButlerEnergySyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaFmhealthButlerEnergySyncAPIResponseModel).Reset()
}

// AlibabaFmhealthButlerEnergySyncAPIResponseModel is 同步用户消耗能量 成功返回结果
type AlibabaFmhealthButlerEnergySyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fmhealth_butler_energy_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaFmhealthButlerEnergySyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaFmhealthButlerEnergySyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaFmhealthButlerEnergySyncAPIResponse)
	},
}

// GetAlibabaFmhealthButlerEnergySyncAPIResponse 从 sync.Pool 获取 AlibabaFmhealthButlerEnergySyncAPIResponse
func GetAlibabaFmhealthButlerEnergySyncAPIResponse() *AlibabaFmhealthButlerEnergySyncAPIResponse {
	return poolAlibabaFmhealthButlerEnergySyncAPIResponse.Get().(*AlibabaFmhealthButlerEnergySyncAPIResponse)
}

// ReleaseAlibabaFmhealthButlerEnergySyncAPIResponse 将 AlibabaFmhealthButlerEnergySyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaFmhealthButlerEnergySyncAPIResponse(v *AlibabaFmhealthButlerEnergySyncAPIResponse) {
	v.Reset()
	poolAlibabaFmhealthButlerEnergySyncAPIResponse.Put(v)
}
