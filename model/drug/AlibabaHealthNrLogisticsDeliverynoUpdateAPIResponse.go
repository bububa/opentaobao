package drug

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahealthnrlogisticsdeliverynoupdateAPIResponse 上传订单同城快递单号 API返回值
// alibaba.health.nr.logistics.deliveryno.update
//
// 上传订单同城快递单号
type AlibabahealthnrlogisticsdeliverynoupdateAPIResponse struct {
	model.CommonResponse
	AlibabahealthnrlogisticsdeliverynoupdateAPIResponseModel
}

// AlibabahealthnrlogisticsdeliverynoupdateAPIResponseModel is 上传订单同城快递单号 成功返回结果
type AlibabahealthnrlogisticsdeliverynoupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_nr_logistics_deliveryno_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 更新成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
