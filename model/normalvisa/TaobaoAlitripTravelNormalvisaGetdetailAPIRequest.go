package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelNormalvisaGetdetailAPIRequest
获取单笔订单的详情 API请求
taobao.alitrip.travel.normalvisa.getdetail

获取单笔签证的详细记录 */
type TaobaoAlitripTravelNormalvisaGetdetailAPIRequest struct {
	model.Params
	// 订单id
	_bizOrderId int64
}

// New
