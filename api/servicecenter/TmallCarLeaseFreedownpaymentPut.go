package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallCarLeaseFreedownpaymentPut 同步直租车免首付商品活动信息
// tmall.car.lease.freedownpayment.put
//
// 汽车行业直租车免首付需求中，用与对商品打标，活动范围设置，在消费者端商品详情页、订单等环节透出，表示该商品为直租免首付商品。
func TmallCarLeaseFreedownpaymentPut(clt *core.SDKClient, req *servicecenter.TmallCarLeaseFreedownpaymentPutAPIRequest, session string) (*servicecenter.TmallCarLeaseFreedownpaymentPutAPIResponse, error) {
	var resp servicecenter.TmallCarLeaseFreedownpaymentPutAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
