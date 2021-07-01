package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTotoroAuxproductDeleteAPIRequest
廉航辅营产品删除 API请求
taobao.alitrip.totoro.auxproduct.delete

廉航辅营产品删除接口 */
type TaobaoAlitripTotoroAuxproductDeleteAPIRequest struct {
	model.Params
	// 廉航辅营产品删除请求
	_delAuxProductRq *DelAuxProductRq
}

// New
