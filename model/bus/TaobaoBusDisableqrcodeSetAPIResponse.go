package bus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusdisableqrcodesetAPIResponse 自助机失效二维码 API返回值
// taobao.bus.disableqrcode.set
//
// 使创建的二维码失效
type TaobaobusdisableqrcodesetAPIResponse struct {
	model.CommonResponse
	TaobaobusdisableqrcodesetAPIResponseModel
}

// TaobaobusdisableqrcodesetAPIResponseModel is 自助机失效二维码 成功返回结果
type TaobaobusdisableqrcodesetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_disableqrcode_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errorCode
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// errorMsg
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
