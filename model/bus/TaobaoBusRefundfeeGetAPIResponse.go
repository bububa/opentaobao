package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询退票费用明细 API返回值 
taobao.bus.refundfee.get

查询退票的费用信息
*/
type TaobaoBusRefundfeeGetAPIResponse struct {
    model.CommonResponse
    TaobaoBusRefundfeeGetAPIResponseModel
}

// 查询退票费用明细 成功返回结果
type TaobaoBusRefundfeeGetAPIResponseModel struct {
    XMLName xml.Name `xml:"bus_refundfee_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *B2BQueryRefundFeeRp `json:"result,omitempty" xml:"result,omitempty"`
}
