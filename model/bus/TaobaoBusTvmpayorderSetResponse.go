package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
自助机条形码被动支付 APIResponse
taobao.bus.tvmpayorder.set

汽车票线下自助机条形码支付
*/
type TaobaoBusTvmpayorderSetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"bus_tvmpayorder_set_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // errorCode  线下扫码支付  错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"