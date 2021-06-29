package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
同城配门店备货通知 APIResponse
tmall.nr.fulfill.logistics.consign

同城配业务备货通知，商家告诉平台门店的货已经准备好，可以发货了；
*/
type TmallNrFulfillLogisticsConsignAPIResponse struct {
    model.CommonResponse
    TmallNrFulfillLogisticsConsignResponse
}

type TmallNrFulfillLogisticsConsignResponse struct {
    XMLName xml.Name `xml:"tmall_nr_fulfill_logistics_consign_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *NrResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
