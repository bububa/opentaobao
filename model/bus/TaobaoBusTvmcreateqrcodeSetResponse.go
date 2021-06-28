package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
自助机生成支付宝支付二维码 APIResponse
taobao.bus.tvmcreateqrcode.set

用于汽车票线下自助机调用获取支付宝的二维码
*/
type TaobaoBusTvmcreateqrcodeSetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"bus_tvmcreateqrcode_set_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // errorCode      错误码 及其 描述 ("PAYED_OR_CANCEL_FAIL_NOT_CREATE_QRCODE", "已支付/取消/失败订单不能创建二维码"),     ("PAYED_OR_CANCEL_FAIL_NOT_CREATE_SCAN_CODE", "已支付/取消/失败订单不能创建条形码"),     ("CREATE_QRCODE_ERROR", "创建二维码失败"),     ("DISABLE_QRCODE_ERROR", "失效二维码失败"),     ("CREATE_SCANCODE_ERROR", "扫码失败"),     ("CREATE_SCANCODE_PROCESSING", "扫码处理中")
    
    ResultCode   string `json:"result_code,omitempty" xml:"