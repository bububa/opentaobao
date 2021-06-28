package tvpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付查询消费抽奖配置 APIResponse
taobao.tvpay.promotion.info.get

查询消费抽奖配置
*/
type TaobaoTvpayPromotionInfoGetAPIResponse struct {
    model.CommonResponse
    TaobaoTvpayPromotionInfoGetResponse
}

type TaobaoTvpayPromotionInfoGetResponse struct {
    XMLName xml.Name `xml:"tvpay_promotion_info_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // Top返回对象
    
    Result   *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
