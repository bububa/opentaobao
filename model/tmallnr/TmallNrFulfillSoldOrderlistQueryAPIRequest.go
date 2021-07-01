package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrFulfillSoldOrderlistQueryAPIRequest
零售商获取品牌商的特定订单列表 API请求
tmall.nr.fulfill.sold.orderlist.query

零售商获取品牌商的特定订单列表，后端服务有零售商和品牌商的绑定关系，存在开关控制；同时后端存在定时送业务等特殊业务的校验，非同城配送业务不能返回，返回值存在品牌方用户的电话号码，当前电话号码是屏蔽中间四位 */
type TmallNrFulfillSoldOrderlistQueryAPIRequest struct {
	model.Params
	// 入参对象
	_param0 *NrTimingOrderSoldQueryReqDto
}

// New
