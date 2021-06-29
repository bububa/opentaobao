package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单笔全渠道经销商订单的详细信息 APIResponse
taobao.omni.dealer.oders.get

全渠道经销商获取单笔订单的详细信息
*/
type TaobaoOmniDealerOdersGetAPIResponse struct {
    model.CommonResponse
    TaobaoOmniDealerOdersGetResponse
}

type TaobaoOmniDealerOdersGetResponse struct {
    XMLName xml.Name `xml:"omni_dealer_oders_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 订单
    
    Data   *TaobaoOmniDealerOdersGetData `json:"data,omitempty" xml:"data,omitempty"`

    
}
