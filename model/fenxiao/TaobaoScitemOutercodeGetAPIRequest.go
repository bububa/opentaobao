package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoScitemOutercodeGetAPIRequest
根据outerCode查询商品 API请求
taobao.scitem.outercode.get

根据outerCode查询商品 */
type TaobaoScitemOutercodeGetAPIRequest struct {
	model.Params
	// 商品编码
	_outerCode string
}

// New
