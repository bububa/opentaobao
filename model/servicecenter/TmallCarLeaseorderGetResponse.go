package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取租赁订单信息 APIResponse
tmall.car.leaseorder.get

卖家通过供销平台获取代销商的订单信息，但是部分情况下网商银行订单号获取不到，需要提供接口或者工具给卖家
*/
type TmallCarLeaseorderGetAPIResponse struct {
    model.CommonResponse
    Response *TmallCarLeaseorderGetResponse `json:"tmall_car_leaseorder_get_response,omitempty"`
}

type TmallCarLeaseorderGetResponse struct {

    // 结果
    Result   *ResultVo `json:"result,omitempty"`

}
