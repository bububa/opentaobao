package vaccin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthvaccinetradesubscribedetailsaveAPIResponse 私立疫苗交易-预约详情更新或保存 API返回值
// alibaba.alihealth.vaccine.trade.subscribe.detail.save
//
// 私立疫苗交易-预约详情更新或保存
type AlibabaalihealthvaccinetradesubscribedetailsaveAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthvaccinetradesubscribedetailsaveAPIResponseModel
}

// AlibabaalihealthvaccinetradesubscribedetailsaveAPIResponseModel is 私立疫苗交易-预约详情更新或保存 成功返回结果
type AlibabaalihealthvaccinetradesubscribedetailsaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_vaccine_trade_subscribe_detail_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务响应code
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 业务错误信息描述
	BizMessage string `json:"biz_message,omitempty" xml:"biz_message,omitempty"`
	// 响应结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 业务成功状态
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
}
