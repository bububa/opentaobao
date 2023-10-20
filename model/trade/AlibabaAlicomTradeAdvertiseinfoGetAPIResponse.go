package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlicomTradeAdvertiseinfoGetAPIResponse 获取订单上的在信息流投放信息 API返回值
// alibaba.alicom.trade.advertiseinfo.get
//
// 获取订单上的在信息流投放信息
type AlibabaAlicomTradeAdvertiseinfoGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlicomTradeAdvertiseinfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlicomTradeAdvertiseinfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlicomTradeAdvertiseinfoGetAPIResponseModel).Reset()
}

// AlibabaAlicomTradeAdvertiseinfoGetAPIResponseModel is 获取订单上的在信息流投放信息 成功返回结果
type AlibabaAlicomTradeAdvertiseinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alicom_trade_advertiseinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 信息投放广告信息
	Module *AdvertiseInfoDto `json:"module,omitempty" xml:"module,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlicomTradeAdvertiseinfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Module = nil
}

var poolAlibabaAlicomTradeAdvertiseinfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlicomTradeAdvertiseinfoGetAPIResponse)
	},
}

// GetAlibabaAlicomTradeAdvertiseinfoGetAPIResponse 从 sync.Pool 获取 AlibabaAlicomTradeAdvertiseinfoGetAPIResponse
func GetAlibabaAlicomTradeAdvertiseinfoGetAPIResponse() *AlibabaAlicomTradeAdvertiseinfoGetAPIResponse {
	return poolAlibabaAlicomTradeAdvertiseinfoGetAPIResponse.Get().(*AlibabaAlicomTradeAdvertiseinfoGetAPIResponse)
}

// ReleaseAlibabaAlicomTradeAdvertiseinfoGetAPIResponse 将 AlibabaAlicomTradeAdvertiseinfoGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlicomTradeAdvertiseinfoGetAPIResponse(v *AlibabaAlicomTradeAdvertiseinfoGetAPIResponse) {
	v.Reset()
	poolAlibabaAlicomTradeAdvertiseinfoGetAPIResponse.Put(v)
}
