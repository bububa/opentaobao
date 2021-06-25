package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分销商提交受理订单 APIResponse
alibaba.tianji.distributor.order.submit

分销商提交受理订单，如合约订购、充值受理等
*/
type AlibabaTianjiDistributorOrderSubmitAPIResponse struct {
    model.CommonResponse
    Response *AlibabaTianjiDistributorOrderSubmitResponse `json:"alibaba_tianji_distributor_order_submit_response,omitempty"`
}

type AlibabaTianjiDistributorOrderSubmitResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
