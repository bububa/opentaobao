package tvpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付申请设备授权 APIResponse
taobao.tvpay.auth.apply

为用户在指定设备上申请支付授权
*/
type TaobaoTvpayAuthApplyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tvpay_auth_apply_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // Top返回对象
    
    Result   *TopResultDo `json:"result,omitempty" xml:"