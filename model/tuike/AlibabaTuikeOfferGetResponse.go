package tuike

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推广商品查询接口 APIResponse
alibaba.tuike.offer.get

查询1688推客平台卖家推广中的商品信息
*/
type AlibabaTuikeOfferGetAPIResponse struct {
    model.CommonResponse
    AlibabaTuikeOfferGetResponse
}

type AlibabaTuikeOfferGetResponse struct {
    XMLName xml.Name `xml:"alibaba_tuike_offer_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询结果模型
    
    Result   *TaOfferSearchResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
