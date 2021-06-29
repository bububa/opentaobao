package tuike

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推广商品查询接口2.0 API返回值 
alibaba.tuike.offer.get.pro

查询1688推客平台卖家推广中的商品信息
*/
type AlibabaTuikeOfferGetProAPIResponse struct {
    model.CommonResponse
    AlibabaTuikeOfferGetProResponse
}

// 推广商品查询接口2.0 成功返回结果
type AlibabaTuikeOfferGetProResponse struct {
    XMLName xml.Name `xml:"alibaba_tuike_offer_get_pro_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 查询结果模型
    Result   *TaOfferSearchResult `json:"result,omitempty" xml:"result,omitempty"`
}
