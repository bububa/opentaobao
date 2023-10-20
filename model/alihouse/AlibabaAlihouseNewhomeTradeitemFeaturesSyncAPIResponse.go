package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhometradeitemfeaturessyncAPIResponse 同步品活动标 API返回值
// alibaba.alihouse.newhome.tradeitem.features.sync
//
// 同步品活动标
type AlibabaalihousenewhometradeitemfeaturessyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhometradeitemfeaturessyncAPIResponseModel
}

// AlibabaalihousenewhometradeitemfeaturessyncAPIResponseModel is 同步品活动标 成功返回结果
type AlibabaalihousenewhometradeitemfeaturessyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_tradeitem_features_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhometradeitemfeaturessyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
