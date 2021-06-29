package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取同城配送业务单笔订单 APIResponse
tmall.nr.order.query.jst

同城配送业务获取单笔订单
*/
type TmallNrOrderQueryJstAPIResponse struct {
    model.CommonResponse
    TmallNrOrderQueryJstResponse
}

type TmallNrOrderQueryJstResponse struct {
    XMLName xml.Name `xml:"tmall_nr_order_query_jst_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *NewRetailResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
