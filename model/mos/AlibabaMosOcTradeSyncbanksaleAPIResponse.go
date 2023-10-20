package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosOcTradeSyncbanksaleAPIResponse 云闪付、银行卡销售数据回传接口 API返回值
// alibaba.mos.oc.trade.syncbanksale
//
// 云闪付、银行卡销售数据回传
type AlibabaMosOcTradeSyncbanksaleAPIResponse struct {
	model.CommonResponse
	AlibabaMosOcTradeSyncbanksaleAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosOcTradeSyncbanksaleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosOcTradeSyncbanksaleAPIResponseModel).Reset()
}

// AlibabaMosOcTradeSyncbanksaleAPIResponseModel is 云闪付、银行卡销售数据回传接口 成功返回结果
type AlibabaMosOcTradeSyncbanksaleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_oc_trade_syncbanksale_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// resultDTO
	Result *AlibabaMosOcTradeSyncbanksaleResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosOcTradeSyncbanksaleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMosOcTradeSyncbanksaleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosOcTradeSyncbanksaleAPIResponse)
	},
}

// GetAlibabaMosOcTradeSyncbanksaleAPIResponse 从 sync.Pool 获取 AlibabaMosOcTradeSyncbanksaleAPIResponse
func GetAlibabaMosOcTradeSyncbanksaleAPIResponse() *AlibabaMosOcTradeSyncbanksaleAPIResponse {
	return poolAlibabaMosOcTradeSyncbanksaleAPIResponse.Get().(*AlibabaMosOcTradeSyncbanksaleAPIResponse)
}

// ReleaseAlibabaMosOcTradeSyncbanksaleAPIResponse 将 AlibabaMosOcTradeSyncbanksaleAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosOcTradeSyncbanksaleAPIResponse(v *AlibabaMosOcTradeSyncbanksaleAPIResponse) {
	v.Reset()
	poolAlibabaMosOcTradeSyncbanksaleAPIResponse.Put(v)
}
