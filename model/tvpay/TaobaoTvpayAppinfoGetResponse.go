package tvpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付获取应用信息 APIResponse
taobao.tvpay.appinfo.get

tv支付获取应用信息
*/
type TaobaoTvpayAppinfoGetAPIResponse struct {
    model.CommonResponse
    TaobaoTvpayAppinfoGetResponse
}

type TaobaoTvpayAppinfoGetResponse struct {
    XMLName xml.Name `xml:"tvpay_appinfo_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // Top返回对象
    
    Result   *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
