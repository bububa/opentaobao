package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询退票费用明细 APIResponse
taobao.bus.refundfee.get

查询退票的费用信息
*/
type TaobaoBusRefundfeeGetAPIResponse struct {
    model.CommonResponse
    TaobaoBusRefundfeeGetResponse
}

type TaobaoBusRefundfeeGetResponse struct {
    XMLName xml.Name `xml:"bus_refundfee_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *B2BQueryRefundFeeRp `json:"result,omitempty" xml:"result,omitempty"`

    
}
