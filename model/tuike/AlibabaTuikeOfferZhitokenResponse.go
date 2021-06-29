package tuike

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
生成阿里口令 APIResponse
alibaba.tuike.offer.zhitoken

推荐链接生产吱口令
*/
type AlibabaTuikeOfferZhitokenAPIResponse struct {
    model.CommonResponse
    AlibabaTuikeOfferZhitokenResponse
}

type AlibabaTuikeOfferZhitokenResponse struct {
    XMLName xml.Name `xml:"alibaba_tuike_offer_zhitoken_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AlibabaTuikeOfferZhitokenResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
