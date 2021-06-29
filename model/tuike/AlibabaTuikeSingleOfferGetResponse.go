package tuike

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询1688推客平台卖家推广中的商品信息 APIResponse
alibaba.tuike.single.offer.get

查询单个推客商品信息的接口
*/
type AlibabaTuikeSingleOfferGetAPIResponse struct {
    model.CommonResponse
    AlibabaTuikeSingleOfferGetResponse
}

type AlibabaTuikeSingleOfferGetResponse struct {
    XMLName xml.Name `xml:"alibaba_tuike_single_offer_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
