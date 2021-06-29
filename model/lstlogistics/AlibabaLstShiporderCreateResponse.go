package lstlogistics

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
零售通发货单创建 APIResponse
alibaba.lst.shiporder.create

通过该接口可以创建零售通运保保发货单，并处理相关业务流程。
*/
type AlibabaLstShiporderCreateAPIResponse struct {
    model.CommonResponse
    AlibabaLstShiporderCreateResponse
}

type AlibabaLstShiporderCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_shiporder_create_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
