package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelNormalvisaGetAPIRequest
获取签证记录 API请求
taobao.alitrip.travel.normalvisa.get

用于获取普通签证的记录信息 */
type TaobaoAlitripTravelNormalvisaGetAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId int64
}

// New
