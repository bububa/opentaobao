package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
店铺收藏 API返回值 
alibaba.interact.shop.favor

店铺收藏mtop接口开放鉴权接口，无入参出参，无安全风险，mtop接口开发 酒仙。
*/
type AlibabaInteractShopFavorAPIResponse struct {
    model.CommonResponse
    AlibabaInteractShopFavorResponse
}

// 店铺收藏 成功返回结果
type AlibabaInteractShopFavorResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_shop_favor_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result=0
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
