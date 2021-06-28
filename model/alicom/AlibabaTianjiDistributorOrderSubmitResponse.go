package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销商提交受理订单 APIResponse
alibaba.tianji.distributor.order.submit

分销商提交受理订单，如合约订购、充值受理等
*/
type AlibabaTianjiDistributorOrderSubmitAPIResponse struct {
    model.CommonResponse
    AlibabaTianjiDistributorOrderSubmitResponse
}

type AlibabaTianjiDistributorOrderSubmitResponse struct {
    XMLName xml.Name `xml:"alibaba_tianji_distributor_order_submit_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
