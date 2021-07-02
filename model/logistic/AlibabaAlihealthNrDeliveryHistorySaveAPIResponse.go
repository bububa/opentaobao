package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthNrDeliveryHistorySaveAPIResponse 物流信息回传接口 API返回值
// alibaba.alihealth.nr.delivery.history.save
//
// 商家ERP回传物流信息
type AlibabaAlihealthNrDeliveryHistorySaveAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthNrDeliveryHistorySaveAPIResponseModel
}

// AlibabaAlihealthNrDeliveryHistorySaveAPIResponseModel is 物流信息回传接口 成功返回结果
type AlibabaAlihealthNrDeliveryHistorySaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_nr_delivery_history_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ResponseResult `json:"result,omitempty" xml:"result,omitempty"`
}
