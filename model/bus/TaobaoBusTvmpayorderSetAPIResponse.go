package bus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusTvmpayorderSetAPIResponse 自助机条形码被动支付 API返回值
// taobao.bus.tvmpayorder.set
//
// 汽车票线下自助机条形码支付
type TaobaoBusTvmpayorderSetAPIResponse struct {
	model.CommonResponse
	TaobaoBusTvmpayorderSetAPIResponseModel
}

// TaobaoBusTvmpayorderSetAPIResponseModel is 自助机条形码被动支付 成功返回结果
type TaobaoBusTvmpayorderSetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_tvmpayorder_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errorCode  线下扫码支付  错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// errorMsg 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// payTime
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// success true 成功 false 失败
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
