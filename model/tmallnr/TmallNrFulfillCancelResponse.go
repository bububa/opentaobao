package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消上门揽件 APIResponse
tmall.nr.fulfill.cancel

新零售门店业务取消上门揽件
*/
type TmallNrFulfillCancelAPIResponse struct {
    model.CommonResponse
    TmallNrFulfillCancelResponse
}

type TmallNrFulfillCancelResponse struct {
    XMLName xml.Name `xml:"tmall_nr_fulfill_cancel_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *NrResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
