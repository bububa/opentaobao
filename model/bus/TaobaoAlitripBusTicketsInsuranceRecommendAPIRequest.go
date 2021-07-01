package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest
汽车票保险推荐 API请求
taobao.alitrip.bus.tickets.insurance.recommend

获取推荐保险内容 */
type TaobaoAlitripBusTicketsInsuranceRecommendAPIRequest struct {
	model.Params
	// 请求对象
	_recommendReq *TopStandardInsRecommendRequest
}

// New
