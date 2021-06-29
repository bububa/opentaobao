package tvpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付 API返回值 
taobao.tvpay.access.data.get

在匿名用户支付后尝试为其登陆绑定的淘宝账号
*/
type TaobaoTvpayAccessDataGetAPIResponse struct {
    model.CommonResponse
    TaobaoTvpayAccessDataGetResponse
}

// tv支付 成功返回结果
type TaobaoTvpayAccessDataGetResponse struct {
    XMLName xml.Name `xml:"tvpay_access_data_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // Top返回对象
    Result   *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
