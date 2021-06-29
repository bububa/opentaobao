package lstlogistics

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-异云-使用第三方物流发货 APIResponse
alibaba.lst.logistics.thirdpart.send

异地云仓的订单，使用第三方物流的发货方式来变更订单发货状态
*/
type AlibabaLstLogisticsThirdpartSendAPIResponse struct {
    model.CommonResponse
    AlibabaLstLogisticsThirdpartSendResponse
}

type AlibabaLstLogisticsThirdpartSendResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_logistics_thirdpart_send_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaLstLogisticsThirdpartSendResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
