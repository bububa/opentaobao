package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaophoneorderexternalcreateAPIResponse 数字虚拟话费外放下单接口 API返回值
// taobao.phone.order.external.create
//
// 数字虚拟话费外放下单接口
type TaobaophoneorderexternalcreateAPIResponse struct {
	model.CommonResponse
	TaobaophoneorderexternalcreateAPIResponseModel
}

// TaobaophoneorderexternalcreateAPIResponseModel is 数字虚拟话费外放下单接口 成功返回结果
type TaobaophoneorderexternalcreateAPIResponseModel struct {
	XMLName xml.Name `xml:"phone_order_external_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
