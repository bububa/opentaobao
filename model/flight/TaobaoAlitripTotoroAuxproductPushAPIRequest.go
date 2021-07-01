package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTotoroAuxproductPushAPIRequest
廉航辅营产品投放 API请求
taobao.alitrip.totoro.auxproduct.push

廉航辅营产品投放接口 */
type TaobaoAlitripTotoroAuxproductPushAPIRequest struct {
	model.Params
	// 廉航辅营产品推送请求
	_pushAuxProductsRq *PushAuxProductsRq
}

// New
