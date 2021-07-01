package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTelecomZhimaOrdernotifyCallbackAPIResponse
阿里通信芝麻订单通知 API返回值
alibaba.telecom.zhima.ordernotify.callback

商家通知阿里通信，芝麻订单状态，阿里通信侧进行代扣支付、发货并确认收货 */
type AlibabaTelecomZhimaOrdernotifyCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaTelecomZhimaOrdernotifyCallbackAPIResponseModel
}

// AlibabaTelecomZhimaOrdernotifyCallbackAPIResponseModel is 阿里通信芝麻订单通知 成功返回结果
type AlibabaTelecomZhimaOrdernotifyCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_telecom_zhima_ordernotify_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参对象
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
