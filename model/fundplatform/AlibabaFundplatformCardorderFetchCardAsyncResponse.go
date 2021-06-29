package fundplatform

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
异步批量生成储值卡 APIResponse
alibaba.fundplatform.cardorder.fetch.card.async

外部业务方异步批量生成储值卡的接口。同步只返回接受成功，异步会通知制卡成功
*/
type AlibabaFundplatformCardorderFetchCardAsyncAPIResponse struct {
    model.CommonResponse
    AlibabaFundplatformCardorderFetchCardAsyncResponse
}

type AlibabaFundplatformCardorderFetchCardAsyncResponse struct {
    XMLName xml.Name `xml:"alibaba_fundplatform_cardorder_fetch_card_async_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *CardFetchAsyncResponse `json:"result,omitempty" xml:"result,omitempty"`

    
}
