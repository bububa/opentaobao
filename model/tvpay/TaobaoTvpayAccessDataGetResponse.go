package tvpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付 APIResponse
taobao.tvpay.access.data.get

在匿名用户支付后尝试为其登陆绑定的淘宝账号
*/
type TaobaoTvpayAccessDataGetAPIResponse struct {
    model.CommonResponse
    TaobaoTvpayAccessDataGetResponse
}

type TaobaoTvpayAccessDataGetResponse struct {
    XMLName xml.Name `xml:"tvpay_access_data_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // Top返回对象
    
    Result   *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
