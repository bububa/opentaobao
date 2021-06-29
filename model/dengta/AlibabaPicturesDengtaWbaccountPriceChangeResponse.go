package dengta

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
微博公众号价格变化通知 APIResponse
alibaba.pictures.dengta.wbaccount.price.change

微博公众号推广价格变更通知接口
*/
type AlibabaPicturesDengtaWbaccountPriceChangeAPIResponse struct {
    model.CommonResponse
    AlibabaPicturesDengtaWbaccountPriceChangeResponse
}

type AlibabaPicturesDengtaWbaccountPriceChangeResponse struct {
    XMLName xml.Name `xml:"alibaba_pictures_dengta_wbaccount_price_change_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ApiGeneralResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
