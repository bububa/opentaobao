package servicecenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarLeaseFreedownpaymentPutAPIResponse
同步直租车免首付商品活动信息 API返回值
tmall.car.lease.freedownpayment.put

汽车行业直租车免首付需求中，用与对商品打标，活动范围设置，在消费者端商品详情页、订单等环节透出，表示该商品为直租免首付商品。 */
type TmallCarLeaseFreedownpaymentPutAPIResponse struct {
	model.CommonResponse
	TmallCarLeaseFreedownpaymentPutAPIResponseModel
}

// TmallCarLeaseFreedownpaymentPutAPIResponseModel is 同步直租车免首付商品活动信息 成功返回结果
type TmallCarLeaseFreedownpaymentPutAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_car_lease_freedownpayment_put_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallCarLeaseFreedownpaymentPutResult `json:"result,omitempty" xml:"result,omitempty"`
}
