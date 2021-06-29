package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询取消履约的原因列表 APIResponse
tmall.nr.fulfill.cancel.reason.query

新零售门店业务查询取消上门揽件的原因列表
*/
type TmallNrFulfillCancelReasonQueryAPIResponse struct {
    model.CommonResponse
    TmallNrFulfillCancelReasonQueryResponse
}

type TmallNrFulfillCancelReasonQueryResponse struct {
    XMLName xml.Name `xml:"tmall_nr_fulfill_cancel_reason_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *NrResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
