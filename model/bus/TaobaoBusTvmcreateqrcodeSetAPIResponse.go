package bus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusTvmcreateqrcodeSetAPIResponse
自助机生成支付宝支付二维码 API返回值
taobao.bus.tvmcreateqrcode.set

用于汽车票线下自助机调用获取支付宝的二维码 */
type TaobaoBusTvmcreateqrcodeSetAPIResponse struct {
	model.CommonResponse
	TaobaoBusTvmcreateqrcodeSetAPIResponseModel
}

// TaobaoBusTvmcreateqrcodeSetAPIResponseModel is 自助机生成支付宝支付二维码 成功返回结果
type TaobaoBusTvmcreateqrcodeSetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_tvmcreateqrcode_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// errorCode      错误码 及其 描述 ("PAYED_OR_CANCEL_FAIL_NOT_CREATE_QRCODE", "已支付/取消/失败订单不能创建二维码"),     ("PAYED_OR_CANCEL_FAIL_NOT_CREATE_SCAN_CODE", "已支付/取消/失败订单不能创建条形码"),     ("CREATE_QRCODE_ERROR", "创建二维码失败"),     ("DISABLE_QRCODE_ERROR", "失效二维码失败"),     ("CREATE_SCANCODE_ERROR", "扫码失败"),     ("CREATE_SCANCODE_PROCESSING", "扫码处理中")
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// errorMsg     错误码 及其 描述 ("PAYED_OR_CANCEL_FAIL_NOT_CREATE_QRCODE", "已支付/取消/失败订单不能创建二维码"),     ("PAYED_OR_CANCEL_FAIL_NOT_CREATE_SCAN_CODE", "已支付/取消/失败订单不能创建条形码"),     ("CREATE_QRCODE_ERROR", "创建二维码失败"),     ("DISABLE_QRCODE_ERROR", "失效二维码失败"),     ("CREATE_SCANCODE_ERROR", "扫码失败"),     ("CREATE_SCANCODE_PROCESSING", "扫码处理中")
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// qrCode 二维码连接
	QrCode string `json:"qr_code,omitempty" xml:"qr_code,omitempty"`
	// success true 成功 false 失败
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
