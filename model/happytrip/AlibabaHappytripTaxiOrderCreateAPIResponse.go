package happytrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahappytriptaxiordercreateAPIResponse 用户叫车 API返回值
// alibaba.happytrip.taxi.order.create
//
// 用户根据需要发起叫车请求，在发起请求之前必须事先获得order id.
type AlibabahappytriptaxiordercreateAPIResponse struct {
	model.CommonResponse
	AlibabahappytriptaxiordercreateAPIResponseModel
}

// AlibabahappytriptaxiordercreateAPIResponseModel is 用户叫车 成功返回结果
type AlibabahappytriptaxiordercreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_happytrip_taxi_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 错误码
	Errno int64 `json:"errno,omitempty" xml:"errno,omitempty"`
	// 订单创建结果
	Data *OrderCreateResult `json:"data,omitempty" xml:"data,omitempty"`
}
