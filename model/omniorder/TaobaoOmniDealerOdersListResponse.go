package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
全渠道经销商订单列表 APIResponse
taobao.omni.dealer.oders.list

全渠道经销商订单列表查询
*/
type TaobaoOmniDealerOdersListAPIResponse struct {
    model.CommonResponse
    TaobaoOmniDealerOdersListResponse
}

type TaobaoOmniDealerOdersListResponse struct {
    XMLName xml.Name `xml:"omni_dealer_oders_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 分页结果
    
    Result   *PageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
