package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarLeaseFreedownpaymentPutAPIRequest
同步直租车免首付商品活动信息 API请求
tmall.car.lease.freedownpayment.put

汽车行业直租车免首付需求中，用与对商品打标，活动范围设置，在消费者端商品详情页、订单等环节透出，表示该商品为直租免首付商品。 */
type TmallCarLeaseFreedownpaymentPutAPIRequest struct {
	model.Params
	// 活动预热结束时间:格式：yyyy.MM.dd HH:mm:ss
	_preEndTime string
	// 活动预热开始时间:格式：yyyy.MM.dd HH:mm:ss
	_preStartTime string
	// 商品ID
	_itemId int64
	// 活动时间范围节点(json格式字符串)：<br/> 开始时间(startTime),格式：yyyy.MM.dd HH:mm:ss <br/>  结束时间(endTime),格式：yyyy.MM.dd HH:mm:ss <br/>  名额(amount)
	_timeRangeList string
	// 外部活动ID
	_refActivityId string
}

// New
