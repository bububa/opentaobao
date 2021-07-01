package happytrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTaxiServicestatusGetAPIResponse
供应商服务开通状态 API返回值
alibaba.happytrip.taxi.servicestatus.get

获取服务供应商在每个地区的服务开通状态、支持的车型等 */
type AlibabaHappytripTaxiServicestatusGetAPIResponse struct {
	model.CommonResponse
	AlibabaHappytripTaxiServicestatusGetAPIResponseModel
}

// AlibabaHappytripTaxiServicestatusGetAPIResponseModel is 供应商服务开通状态 成功返回结果
type AlibabaHappytripTaxiServicestatusGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_servicestatus_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 供应商服务状态数据
	Data *ServiceStatusModel `json:"data,omitempty" xml:"data,omitempty"`
	// 错误代码
	Errno int64 `json:"errno,omitempty" xml:"errno,omitempty"`
	// 错误消息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
}
