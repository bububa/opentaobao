package dengta

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
微信公众号价格变化通知 APIResponse
alibaba.pictures.dengta.wxaccount.price.change

微信公众号推广价格变更通知接口
*/
type AlibabaPicturesDengtaWxaccountPriceChangeAPIResponse struct {
    model.CommonResponse
    AlibabaPicturesDengtaWxaccountPriceChangeResponse
}

type AlibabaPicturesDengtaWxaccountPriceChangeResponse struct {
    XMLName xml.Name `xml:"alibaba_pictures_dengta_wxaccount_price_change_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ApiGeneralResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
