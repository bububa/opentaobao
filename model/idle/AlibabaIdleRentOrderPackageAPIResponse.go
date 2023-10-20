package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerentorderpackageAPIResponse 确认揽收商品 API返回值
// alibaba.idle.rent.order.package
//
// 确认揽收
type AlibabaidlerentorderpackageAPIResponse struct {
	model.CommonResponse
	AlibabaidlerentorderpackageAPIResponseModel
}

// AlibabaidlerentorderpackageAPIResponseModel is 确认揽收商品 成功返回结果
type AlibabaidlerentorderpackageAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_rent_order_package_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaidlerentorderpackageTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
