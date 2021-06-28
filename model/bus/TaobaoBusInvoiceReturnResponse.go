package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发票回调接口 APIResponse
taobao.bus.invoice.return

汽车票发票回调接口
*/
type TaobaoBusInvoiceReturnAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"bus_invoice_return_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"