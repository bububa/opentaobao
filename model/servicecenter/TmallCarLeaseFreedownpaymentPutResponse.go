package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
同步直租车免首付商品活动信息 APIResponse
tmall.car.lease.freedownpayment.put

汽车行业直租车免首付需求中，用与对商品打标，活动范围设置，在消费者端商品详情页、订单等环节透出，表示该商品为直租免首付商品。
*/
type TmallCarLeaseFreedownpaymentPutAPIResponse struct {
    model.CommonResponse
    Response *TmallCarLeaseFreedownpaymentPutResponse `json:"tmall_car_lease_freedownpayment_put_response,omitempty"`
}

type TmallCarLeaseFreedownpaymentPutResponse struct {

    // 接口返回model
    Result   *TmallCarLeaseFreedownpaymentPutResult `json:"result,omitempty"`

}
