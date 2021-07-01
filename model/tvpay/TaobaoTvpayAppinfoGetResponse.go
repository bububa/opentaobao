package tvpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付获取应用信息 API返回值 
taobao.tvpay.appinfo.get

tv支付获取应用信息
*/
type TaobaoTvpayAppinfoGetAPIResponse struct {
    model.CommonResponse
    TaobaoTvpayAppinfoGetResponse
}

// tv支付获取应用信息 成功返回结果
type TaobaoTvpayAppinfoGetResponse struct {
    XMLName xml.Name `xml:"tvpay_appinfo_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // Top返回对象
    Result   *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
