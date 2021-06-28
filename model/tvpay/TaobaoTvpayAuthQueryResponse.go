package tvpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付授权查询 APIResponse
taobao.tvpay.auth.query

查询该用户在指定设备上是否有支付授权
*/
type TaobaoTvpayAuthQueryAPIResponse struct {
    model.CommonResponse
    TaobaoTvpayAuthQueryResponse
}

type TaobaoTvpayAuthQueryResponse struct {
    XMLName xml.Name `xml:"tvpay_auth_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // Top返回对象
    
    Result   *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
