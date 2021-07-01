package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest
获取物流公司信息 API请求
taobao.alitrip.travel.normalvisa.getcompany

获取物流公司信息 */
type TaobaoAlitripTravelNormalvisaGetcompanyAPIRequest struct {
	model.Params
	// true：取5个重要的物流公司 false：取所有的物流公司
	_param0 bool
}

// New
