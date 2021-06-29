package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
不落库商家推送更新酒店rate APIResponse
taobao.xhotel.intl.rate.update

商家主动推送不落库商品的酒店信息
*/
type TaobaoXhotelIntlRateUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelIntlRateUpdateResponse
}

type TaobaoXhotelIntlRateUpdateResponse struct {
    XMLName xml.Name `xml:"xhotel_intl_rate_update_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询结果集
    
    Result   *TaobaoXhotelIntlRateUpdateResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
