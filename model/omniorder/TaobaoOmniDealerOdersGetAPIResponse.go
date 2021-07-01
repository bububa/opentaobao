package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单笔全渠道经销商订单的详细信息 API返回值 
taobao.omni.dealer.oders.get

全渠道经销商获取单笔订单的详细信息
*/
type TaobaoOmniDealerOdersGetAPIResponse struct {
    model.CommonResponse
    TaobaoOmniDealerOdersGetAPIResponseModel
}

// 获取单笔全渠道经销商订单的详细信息 成功返回结果
type TaobaoOmniDealerOdersGetAPIResponseModel struct {
    XMLName xml.Name `xml:"omni_dealer_oders_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 订单
    Data   *TaobaoOmniDealerOdersGetData `json:"data,omitempty" xml:"data,omitempty"`
}
