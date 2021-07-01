package tvpay

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付查询消费抽奖配置 API返回值 
taobao.tvpay.promotion.info.get

查询消费抽奖配置
*/
type TaobaoTvpayPromotionInfoGetAPIResponse struct {
    model.CommonResponse
    TaobaoTvpayPromotionInfoGetAPIResponseModel
}

// tv支付查询消费抽奖配置 成功返回结果
type TaobaoTvpayPromotionInfoGetAPIResponseModel struct {
    XMLName xml.Name `xml:"tvpay_promotion_info_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // Top返回对象
    Result   *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
