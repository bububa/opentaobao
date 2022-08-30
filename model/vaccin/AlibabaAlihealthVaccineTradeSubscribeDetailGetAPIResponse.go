package vaccin

import (
	"encoding/xml"

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

// AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponseModel is 私立疫苗交易-预约详情获取 成功返回结果
type AlibabaAlihealthVaccineTradeSubscribeDetailGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_vaccine_trade_subscribe_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 预约详情响应
	Data *TradeVaccineSubscribeDetailTopResult `json:"data,omitempty" xml:"data,omitempty"`
}
