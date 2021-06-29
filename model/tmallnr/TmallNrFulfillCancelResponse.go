package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消上门揽件 API返回值 
tmall.nr.fulfill.cancel

新零售门店业务取消上门揽件
*/
type TmallNrFulfillCancelAPIResponse struct {
    model.CommonResponse
    TmallNrFulfillCancelResponse
}

// 取消上门揽件 成功返回结果
type TmallNrFulfillCancelResponse struct {
    XMLName xml.Name `xml:"tmall_nr_fulfill_cancel_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *NrResult `json:"result,omitempty" xml:"result,omitempty"`
}
