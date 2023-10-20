package vaccin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponse 私立疫苗交易-预约详情获取 API返回值
// alibaba.alihealth.vaccine.trade.subscribe.detail.get
//
// 私立疫苗交易-预约详情获取
type AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponseModel).Reset()
}

// AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponseModel is 私立疫苗交易-预约详情获取 成功返回结果
type AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_vaccine_trade_subscribe_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务报错code
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 业务报错消息
	BizMessage string `json:"biz_message,omitempty" xml:"biz_message,omitempty"`
	// 预约详情响应
	Data *TradeVaccineSubscribeDetailTopResult `json:"data,omitempty" xml:"data,omitempty"`
	// 业务成功状态
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizCode = ""
	m.BizMessage = ""
	m.Data = nil
	m.BizSuccess = false
}

var poolAlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponse)
	},
}

// GetAlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponse 从 sync.Pool 获取 AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponse
func GetAlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponse() *AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponse {
	return poolAlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponse.Get().(*AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponse)
}

// ReleaseAlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponse 将 AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponse(v *AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponse.Put(v)
}
