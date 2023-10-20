package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhometradeitemrelationAPIResponse 货独立绑定货品 API返回值
// alibaba.alihouse.newhome.tradeitem.relation
//
// 货独立绑定货品
type AlibabaalihousenewhometradeitemrelationAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhometradeitemrelationAPIResponseModel
}

// AlibabaalihousenewhometradeitemrelationAPIResponseModel is 货独立绑定货品 成功返回结果
type AlibabaalihousenewhometradeitemrelationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_tradeitem_relation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhometradeitemrelationResult `json:"result,omitempty" xml:"result,omitempty"`
}
