package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询退票费用明细 APIResponse
taobao.bus.refundfee.get

查询退票的费用信息
*/
type TaobaoBusRefundfeeGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBusRefundfeeGetResponse `json:"taobao_bus_refundfee_get_response,omitempty"`
}

type TaobaoBusRefundfeeGetResponse struct {

    // result
    Result   *B2BQueryRefundFeeRp `json:"result,omitempty"`

}
